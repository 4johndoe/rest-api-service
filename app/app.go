package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8080", router)
}
