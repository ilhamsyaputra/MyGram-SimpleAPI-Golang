package helper

import jwt "github.com/dgrijalva/jwt-go"

var secretKet = "rangerhitam142"

func GenerateToken(id string, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKet))

	return signedToken
}
