package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func forwardRequest(w http.ResponseWriter, r *http.Request, url string) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://user_service:8000" + r.URL.Path
	forwardRequest(w, r, url)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	url := "http://product_service:8001" + r.URL.Path
	forwardRequest(w, r, url)
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	url := "http://order_service:8002" + r.URL.Path
	forwardRequest(w, r, url)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handleUsers).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", handleUsers).Methods("GET")

	r.HandleFunc("/products", handleProducts).Methods("GET", "POST")
	r.HandleFunc("/products/{id}", handleProducts).Methods("GET")

	r.HandleFunc("/orders", handleOrders).Methods("GET", "POST")
	r.HandleFunc("/orders/{id}", handleOrders).Methods("GET")

	fmt.Println("Starting API Gateway on port 8080")
	http.ListenAndServe(":8080", r)
}
