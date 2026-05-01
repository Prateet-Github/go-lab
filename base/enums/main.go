package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

func main() {
	var s Status = Pending

	fmt.Println(s)
}
