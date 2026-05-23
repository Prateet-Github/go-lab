package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	user := User{
		Name:  "John",
		Age:   30,
		Email: "prateettiwari29@gmail.com",
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Email)
}
