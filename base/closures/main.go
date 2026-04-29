package main

// Counter returns a closure that keeps track of how many times it has been called.
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
