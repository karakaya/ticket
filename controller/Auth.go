package controller

import (
	"net/http"
	"ticket/config"
	"ticket/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(password) == 0 || len(username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("len 0!"))
		return
	}

	var user models.User
	config.DB.First(&user).Where("username = ?", username)
	check := CheckPasswordHash(password, user.Password)

	if check {
		token, err := getToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating jwt token: " + err.Error()))
			return
		}
		w.Header().Set("Authorization", "Bearer "+token)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorization: " + token))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("credentials do not match"))
}

func getToken(name string) (string, error) {
	singingKey := []byte("mykey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,

		"ExpiresAt": time.Now().Add(time.Minute * 20).Unix(),
	})

	tokenString, err := token.SignedString(singingKey)
	return tokenString, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	singingKey := []byte("mykey")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return singingKey, nil
	})
	errExp := token.Claims.Valid()
	if errExp != nil {
		panic(errExp)
	}

	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
func Protected(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("you are in!"))
}
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
