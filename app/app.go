package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8080", router)
}
