package authentication

import "golang.org/x/crypto/bcrypt"

func ValidateToken() {}

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
