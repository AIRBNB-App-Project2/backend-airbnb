package middlewares

import (
	"be/configs"
	"be/entities"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u entities.User) (string, error) {
	if u.ID == 0 {
		return "cannot Generate token", errors.New("id == 0")
	}

	codes := jwt.MapClaims{
		"user_uid": u.User_uid,
		"email":    u.Email,
		"password": u.Password,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"auth":     true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserUid(e echo.Context) string {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		user_uid := codes["user_uid"].(string)
		return user_uid
	}
	return ""
}

func ExtractTokenAdmin(e echo.Context) (result string) {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		result = codes["email"].(string)
		// result[1] = codes["password"].(string)
		return result
	}
	return ""
}
