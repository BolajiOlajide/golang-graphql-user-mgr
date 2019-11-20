package golang_graphql_user_mgr

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

// JWTDecode method for decoding a JWT token
func JWTDecode(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}

// JWTCreate method for creating a JWT token
func JWTCreate(userID int, expiredAt int64) string {
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    "proton",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(mySigningKey)
	return signedToken
}
