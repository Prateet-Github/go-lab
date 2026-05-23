package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Task", id, "completed")
}

func main() {
	for i := 1; i <= 1000; i++ {
		go worker(i)
	}

	time.Sleep(2 * time.Second)
}
