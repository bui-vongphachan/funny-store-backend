package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/admins"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Login(db *mongo.Database, c *gin.Context) {
	result := gin.H{
		"data":    nil,
		"message": "Invalid data",
	}

	log.Println("API_Login")
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&body)

	admin, err := admins.FindOneByEmail(db, &body.Email)
	if err != nil {
		result["message"] = err.Error()
		c.JSON(http.StatusBadRequest, result)
		return
	}

	isMatched, err := ComparePassword(&admin.Password, &body.Password)
	if err != nil {
		result["message"] = err.Error()
		c.JSON(http.StatusBadRequest, result)
		return
	}

	if !isMatched {
		result["message"] = "Invalid email or password"
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result["data"] = nil
	result["message"] = "Success"

	c.JSON(http.StatusOK, result)
}
