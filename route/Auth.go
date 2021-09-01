package route

import (
	"ticket/controller"
	"ticket/middleware"

	"github.com/gorilla/mux"
)

func Authroutes(a *mux.Router) {

	a.HandleFunc("/login", controller.Login).Methods("POST")
	a.HandleFunc("/login-protected", middleware.IsAuth(controller.Protected)).Methods("GET")

}
