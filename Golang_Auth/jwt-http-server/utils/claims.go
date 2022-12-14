package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// use to verify if user exists in the DB (in our case the map) and if exists then create claim for a user
func CreateClaim(username string) (UserClaims, error) {
	if usernames[username] == "" {
		return UserClaims{}, fmt.Errorf("Invalid User")
	}
	userclaim := UserClaims{
		jwt.StandardClaims{
			Audience:  "client",
			ExpiresAt: time.Now().Add(time.Duration(60) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "sauravrana646",
		},
		username,
	}
	return userclaim,nil
}
