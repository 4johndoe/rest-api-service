package app

import (
	"banking/service"
	"encoding/json"
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

	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
