package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// keys map to impersonate DB of keys 
var Keys map[string][]byte = map[string][]byte{
	"HS256": []byte(`hello`),
	"RSA": []byte(`abc`), // here we have to pass *rsa.publickey
	"ECDSA": []byte(`def`), // here we have to pass *ecdsa.publickey
}

// usernames map to impersonate DB of users. This will be used to check if user who requested the toknen is a valid user in the DB.

var usernames map[string]string = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

// User defined type to store claims for JWT
type UserClaims struct {
	// this is a custom claim that we can use to add additional sections in JWT
	jwt.StandardClaims
	Username string
}

// any custom claim interface should implement valid function to be a part of token interface (user in token.Valid call )
func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}
	if usernames[u.Username] == "" {
		return fmt.Errorf("Invalid Username")
	}
	return nil
}

// Used to generate the JWT token itself
func CreateToken(u *UserClaims) (string, int64, error) {
	
	// this function basically takes the claims and calculates the header for the JWT and then appends the Headers and converts the claims into JSON object and joins them with . (dot) as separator
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)

	// This anonymous function takes token as input and checks the Algorithm used for token and according pass the key if HMAC256 is used then pass the HMAC256's corresponding key from Keys DB map
	signingkey := func (t *jwt.Token) interface{}{
		if t.Method.Alg() == jwt.SigningMethodHS256.Alg(){
			return Keys["HS256"]
		}
		return nil
	}(token)


	// This method generates a signature using the key and the HMAC's Sign function and then appends the signature to the previous token generated in above step. Thereby creating a signed token with all three section of a JWT
	signedToken, err := token.SignedString(signingkey)
	if err != nil {
		return "", token.Claims.(*UserClaims).ExpiresAt,fmt.Errorf("Error in createToken when signing Token : %w", err)
	}

	return signedToken, token.Claims.(*UserClaims).ExpiresAt, nil
}

func ParseToken(token string) (*UserClaims, error) {

	// this function is used to verify the signature using the callback func as third parameter which is used to decide which key to use to verify the signature.
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		// if algo match then we can use our key to sign the token
		if t.Method.Alg() == jwt.SigningMethodHS256.Alg() {
			return Keys["HS256"], nil
		}
		return nil, fmt.Errorf("Invalid Signing Algorithm")
		
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken : %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, Invalid Token")
	}

	// returning custom claims if token signature is verified and parsed
	return t.Claims.(*UserClaims), nil
}
