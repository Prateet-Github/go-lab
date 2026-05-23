package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
)

func main() {
	var s Status = Pending

	fmt.Println(s)
}