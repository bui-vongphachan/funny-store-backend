package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/admins"
	log_file "github.com/vongphachan/funny-store-backend/src/modules/log-file"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route_Login(db *mongo.Database, c *gin.Context) {
	result := gin.H{
		"data":    nil,
		"message": "Invalid data",
	}

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.Bind(&body)
	if err != nil {
		log_file.LogErrorAndResponse(c, &err, &result, http.StatusBadRequest)
		return
	}

	admin, err := admins.FindOneByEmail(db, &body.Email)
	if err != nil {
		result["message"] = "Invalid email or password"
		log_file.LogErrorAndResponse(c, &err, &result, http.StatusBadRequest)
		return
	}

	isMatched, err := ComparePassword(&admin.Password, &body.Password)
	if err != nil {
		result["message"] = "Invalid email or password"
		log_file.LogErrorAndResponse(c, &err, &result, http.StatusBadRequest)
		return
	}

	if !isMatched {
		result["message"] = "Invalid email or password"
		log_file.LogErrorAndResponse(c, &err, &result, http.StatusBadRequest)
		return
	}

	result["data"] = nil
	result["message"] = "Success"

	log_file.SaveResponseLog(c, &result)
	c.JSON(http.StatusOK, result)
}
