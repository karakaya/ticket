package ticket

import (
	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}

}

type resource struct {
	ser Service
}

//TODO add endpoints
