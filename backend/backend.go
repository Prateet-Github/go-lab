package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users")
	})

	http.HandleFunc("/users/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/profile")
	})

	http.HandleFunc("/users/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/settings")
	})

	http.HandleFunc("/users/latest-orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/latest-orders")
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend received:", r.URL.Path)
	})

	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "FILES BACKEND")
	})

	http.HandleFunc("/users/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/orders")
	})

	http.HandleFunc("/users/42", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/42")
	})

	http.HandleFunc("/users/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /users/info")
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /orders")
	})

	http.HandleFunc("/orders/history", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /orders/history")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /products")
	})

	http.HandleFunc("/products/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET /products/latest")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("Users service listening on :9000")
	http.ListenAndServe(":9000", nil)
}
