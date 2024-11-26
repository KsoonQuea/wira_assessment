package services

import (
	"database/sql"
	"os"
	"time"

	"wira-backend/connection"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(username, password string) (string, error) {
	var hashedPassword string
	err := connection.DB.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}