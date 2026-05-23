package main

import "fmt"

func counter() func() int {
	x := 0 // closure captures x and retains access to it even after counter() has finished executing since it was on stack

	return func() int {
		x++ // heap allocated variable, so it is not destroyed after counter() finishes executing
		return x
	}
}

func main() {

	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
