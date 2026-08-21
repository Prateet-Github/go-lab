package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users from backend :%s\n", port)
	})

	mux.HandleFunc("/users/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/profile from backend :%s\n", port)
	})

	mux.HandleFunc("/users/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/settings from backend :%s\n", port)
	})
	mux.HandleFunc("/users/settings/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/settings/update from backend :%s\n", port)
	})

	mux.HandleFunc("/users/latest-orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/latest-orders from backend :%s\n", port)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend received: %s from backend :%s\n", r.URL.Path, port)
	})

	mux.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "FILES BACKEND from backend :%s\n", port)
	})

	mux.HandleFunc("/users/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/orders from backend :%s\n", port)
	})

	mux.HandleFunc("/users/42", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/42 from backend :%s\n", port)
	})

	mux.HandleFunc("/users/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /users/info from backend :%s\n", port)
	})

	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /orders from backend :%s\n", port)
	})

	mux.HandleFunc("/orders/history", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /orders/history from backend :%s\n", port)
	})

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /products from backend :%s\n", port)
	})

	mux.HandleFunc("/products/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /products/latest from backend :%s\n", port)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /health from backend :%s\n", port)
	})

	log.Printf("Users service listening on :%s", port)

	log.Fatal(
		http.ListenAndServe(":"+port, mux),
	)
}
