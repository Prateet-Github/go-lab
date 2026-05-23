package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("Hello Go")
}

// will cover rest file handling while developing a real project
