package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const secret="rajapremsai"

func GenerateToken(email string , id int)(string, error){
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email": email,
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	})

	tokenStr,err := token.SignedString([]byte(secret))

	return tokenStr,err
}

func TokenCheck(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		id := int(claims["id"].(float64))
		println("Email:", email, "ID:", id)
	} else {
		return jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}

	return nil
}