package route

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	Authroutes(r)
	Ticket(r)
	User(r)
	Category(r)
}
