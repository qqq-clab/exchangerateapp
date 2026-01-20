package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err

}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedtoken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedtoken, err
}

func CheckPassword(password string, hashpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(password))
	return err == nil
}

func ParseJWT(tokenstring string) (string, error) {
	if len(tokenstring) > 7 && tokenstring[:7] == "Bearer " {
		tokenstring = tokenstring[7:]
	}

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpexted signing method")

		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username is not a string")

		}
		return username, nil
	}

	return "", err

}
