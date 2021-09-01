package middleware

import (
	"net/http"
	"ticket/controller"

	"github.com/dgrijalva/jwt-go"
)

func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing auth header"))
			return
		}

		claims, err := controller.VerifyToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying jwt token: " + err.Error()))
			return
		}
		username := claims.(jwt.MapClaims)["username"].(string)
		r.Header.Set("username", username)
		next.ServeHTTP(w, r)
	})
}

/*
func RedirectIfAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) != 0 {
			w.WriteHeader(http.StatusContinue)
			w.Write([]byte("redirect to /"))
			return
		}

	})
}
*/
