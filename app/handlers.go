package app

import (
	"banking/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Vasa", "Kyiv", "123455"},
	//	{"Petya", "Dnipro", "49583"},
	//}

	customers, err := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	//customers := []Customer{
	//	{"Vasa", "Kyiv", "123455"},
	//	{"Petya", "Dnipro", "49583"},
	//}

	customer, err := ch.service.GetCustomer(id)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}
