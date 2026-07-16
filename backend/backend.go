package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from users-service (microservice)!")
	})

	http.ListenAndServe(":9000", nil)
}
