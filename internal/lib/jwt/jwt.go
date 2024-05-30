package jwt

import (
	"diplomServer/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

const Secret = "rtbrtb"

func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
