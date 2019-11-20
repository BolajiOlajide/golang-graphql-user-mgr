package golang_graphql_user_mgr

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// UserAuth structure for auth details
type UserAuth struct {
	UserID    int
	Roles     []string
	IPAddress string
	Token     string
}

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// UserClaims structure for claims representing a user
type UserClaims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

// Middleware middleware
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := TokenFromHTTPRequest(r)

			userID := UserIDFromToken(token)

			if userID == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				io.WriteString(w, `{"error":"invalid_credentials"}`)
				return
			}

			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			userAuth := UserAuth{
				UserID:    userID,
				IPAddress: ip,
				Token:     token,
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// TokenFromHTTPRequest grab token from a request
func TokenFromHTTPRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")

	var tokenString string
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}

	return tokenString
}

// UserIDFromToken decode the userID from a token
func UserIDFromToken(tokenString string) int {
	token, err := JWTDecode(tokenString)
	if err != nil {
		return 0
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims == nil {
			return 0
		}
		return claims.UserID
	}
	return 0
}

// ForContext I don't know what this does, YET!
func ForContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	if raw == nil {
		return nil
	}
	return raw.(*UserAuth)
}

// GetAuthFromContext Oh it makes sense now
func GetAuthFromContext(ctx context.Context) *UserAuth {
	return ForContext(ctx)
}
