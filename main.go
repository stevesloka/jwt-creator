package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	issuer := "ef3e6aa3ff61479996037338df4c0a2e"
	secret := "2df50753b7a3499fb6df71e6ed8432d6"

	var mySigningKey = []byte(secret)

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["iss"] = issuer
	claims["admin"] = true
	claims["name"] = "Steve Sloka"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)
	fmt.Printf(tokenString)
}
