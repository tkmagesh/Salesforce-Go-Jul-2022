package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Cost  float32 `json:"cost"`
	Units int     `json:"units"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Web Server!"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("All the products will be listed"))
	case "POST":
		decoder := json.NewDecoder(r.Body)
		product := new(Product)
		err := decoder.Decode(product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(product)
		product.Id = product.Id + 1
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(product); err != nil {
			fmt.Println("Serialization error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "DELETE":
		w.Write([]byte("The given product will be removed"))
	case "PUT":
		w.Write([]byte("The given product will be updated"))
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}

/*
	handler functions
		determine the request method (GET, POST, PUT etc)
		extract data from route parameter (http://localhost:8080/products/101)
		if req contains payload
			deserialise the payload (encoding/json package)
		process the request
		determine result to be sent to the client
		serialize the result
		send the response
*/
