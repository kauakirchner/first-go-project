package model

import (
	"fmt"
	"os"
	"time"

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
