package utils

import (

	"movie_storage/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtSecretKey = os.Getenv("CONFIG_PATH") 

func ValidatePayloadToken(body models.LoginRequest, claims jwt.MapClaims) bool {
    claimsUsername, okUsername := claims["username"].(string)
    claimsPassword, okPassword := claims["password"].(string)

    if !okUsername || !okPassword || claimsUsername != body.Username || claimsPassword != body.Password {
        return false
    }

    return true
}

func ValidateToken(body models.LoginRequest, claims jwt.MapClaims) bool {
    _, okUsername := claims["username"].(string)

    _, okPassword := claims["password"].(string)

    if !okUsername || !okPassword{
        return false
    }
    return true
}

func CreateAccessToken(username string, password string) (string, error) {
	expTime := time.Now().Add(1 * time.Hour).Unix()
	payload := jwt.MapClaims{
		"username":     username,
		"password": 	password,
		"exp":      	expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}


func ParseAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidType
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}
