package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8080", router)
}
