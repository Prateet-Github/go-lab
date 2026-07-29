package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(30 * time.Second) // Wait for the goroutine to finish
}

func sayHello() {
	fmt.Println("Hello, World")
}
