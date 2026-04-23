package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Let's Go!")
		})

	fmt.Println("Server is runnig on port 80")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
