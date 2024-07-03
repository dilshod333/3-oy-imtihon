package main

import (
	"api-gateway/api/handler"
	"log"
	"net/http"
)

 

func main() {
	r := http.NewServeMux()

	r.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().NewUser(w, r)
	})

	r.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().GetUserByID(w, r)
	})

	r.HandleFunc("PUT /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().UpdateUserr(w, r)
	})

	r.HandleFunc("DELETE /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().DeleteUserr(w, r)
	})


	r.HandleFunc("GET /product/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().GetProductt(w, r)
	})


	r.HandleFunc("POST /product", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().CreateProduct(w, r)
	})

	r.HandleFunc("PUT /product", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().UpdateProductt(w, r)
	})

	r.HandleFunc("DELETE /product/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().DeleteProductt(w, r)
	})

	r.HandleFunc("POST /order", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().CreateOrderr(w, r)
	})

	r.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		handler.Conn().CreateOrders(w, r)
	})

	log.Println("Api-Gateway running on :9001")
	http.ListenAndServe(":9001", r)
}