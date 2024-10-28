package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Order struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

var orders []Order

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	orderID, _ := strconv.Atoi(params["id"])
	for _, order := range orders {
		if order.ID == orderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = len(orders) + 1
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func main() {
	r := mux.NewRouter()

	orders = append(orders, Order{ID: 1, UserID: 1, ProductID: 1, Quantity: 2})
	orders = append(orders, Order{ID: 2, UserID: 2, ProductID: 2, Quantity: 1})

	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")

	fmt.Println("Starting Order Service on port 8002")
	http.ListenAndServe(":8002", r)
}
