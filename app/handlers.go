package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Vasa", "Kyiv", "123455"},
		{"Petya", "Dnipro", "49583"},
	}

	fmt.Println(r.Header)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
