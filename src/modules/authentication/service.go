package authentication

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vongphachan/funny-store-backend/src/modules/admins"
	log_file "github.com/vongphachan/funny-store-backend/src/modules/log-file"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(c *gin.Context) {

}

func GenerateToken(c *gin.Context, admin *admins.Admin) *string {
	secretKey := []byte(KEY_Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": admin,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log_file.SaveErrorLog(c, &err)
		return nil
	}

	return &tokenString
}

func HashPassword(password *string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(input *string, password *string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(*input), []byte(*password))
	if err != nil {
		return false, err
	}

	return true, nil
}
