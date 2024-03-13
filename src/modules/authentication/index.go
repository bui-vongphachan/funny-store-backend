package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/admins"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Login(db *mongo.Database, r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
		}

		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		c.BindJSON(&body)

		admin, err := admins.FindOneByEmail(db, &body.Email)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		isMatched, err := ComparePassword(&admin.Password, &body.Password)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		if !isMatched {
			result["message"] = "Invalid email or password"
			c.JSON(http.StatusOK, result)
			return
		}

		result["data"] = nil
		result["status"] = http.StatusOK
		result["isError"] = false
		result["message"] = "Success"

		c.JSON(http.StatusOK, result)
	})
}
