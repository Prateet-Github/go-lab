package main

import (
	"fmt"
	"maps"
)

func main() {
	user1 := map[string]int{
		"age": 21,
	}

	// fmt.Println(user)
	// fmt.Println(user["age"])

	user2 := make(map[string]int)

	user2["age"] = 21
	// user2["bday"] = 29
	// fmt.Println(user)

	// delete(user, "bday")

	// clear(user)

	v, ok := user2["age"]

	fmt.Println(v)

	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not")
	}

	fmt.Println(maps.Equal(user1, user2))

}
