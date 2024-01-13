package utils

import (
	"fmt"
	"strconv"
	"student_classes_management_service/pkg/application/constant"
	"student_classes_management_service/pkg/application/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	ID       string    `json:"id"`
	UserName string `json:"userName"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.Users) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		ID:       strconv.Itoa(user.UserId),
		UserName: user.Username,
		Name:     user.FullName,
		Role:     user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Token expires in 1 hour
		},
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	return jwtToken
}

func ParseAndValidateToken(jwtToken string) (*jwt.Token, error) {
	return jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
}

func GetTokenClaims(e echo.Context) (*model.AuthClaims, error) {
	rawClaims := e.Get(constant.Claims)
	claims, ok := rawClaims.(*model.AuthClaims)

	if !ok {
		return nil, echo.ErrForbidden
	}

	return claims, nil
}
