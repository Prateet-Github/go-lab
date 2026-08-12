package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users")
	})

	mux.HandleFunc("/users/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/profile")
	})

	mux.HandleFunc("/users/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/settings")
	})

	mux.HandleFunc("/users/latest-orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/latest-orders")
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend received:", r.URL.Path)
	})

	mux.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "FILES BACKEND")
	})

	mux.HandleFunc("/users/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/orders")
	})

	mux.HandleFunc("/users/42", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/42")
	})

	mux.HandleFunc("/users/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/info")
	})

	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /orders")
	})

	mux.HandleFunc("/orders/history", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /orders/history")
	})

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /products")
	})

	mux.HandleFunc("/products/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /products/latest")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("Users service listening on :%s", port)

	log.Fatal(
		http.ListenAndServe(":"+port, mux),
	)
}
