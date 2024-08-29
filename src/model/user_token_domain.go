package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"go.uber.org/zap"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	logger.Info("Init GenerateToken model", zap.String("journey", "loginUser"))
	secretKey := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		logger.Error(
			"Error trying to generate token",
			err,
			zap.String("journey", "loginUser"),
		)
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err%s", err.Error()),
		)
	}
	logger.Info("Token generated succesfully", zap.String("journey", "loginUser"))
	return tokenString, nil
}

func AuthMiddleware(
	ctx *gin.Context,
) {

	secretKey := os.Getenv(JWT_SECRET_KEY)
	token := RemoveBearerPrefix(
		ctx.Request.Header.Get("Authorization"),
	)

	tokenValue, err := jwt.Parse(RemoveBearerPrefix(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		ctx.JSON(errRest.Code, errRest)
		ctx.Abort()
		return
	}

	claims, ok := tokenValue.Claims.(jwt.MapClaims)

	if !ok || !tokenValue.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		ctx.JSON(errRest.Code, errRest)
		ctx.Abort()
		return
	}

	userDomain := userDomain{
		ID:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}

	logger.Info(
		fmt.Sprintf("User authenticated: %#v", userDomain),
	)
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.Trim("Bearer ", token)
	}

	return token
}
