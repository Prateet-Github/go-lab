package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "Prateet",
		Age:  21,
	}

	fmt.Println(u)
}