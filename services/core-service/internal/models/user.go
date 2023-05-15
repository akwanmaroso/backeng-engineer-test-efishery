package models

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	jwt.StandardClaims
	User struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Role        string `json:"role"`
	} `json:"user"`
}
