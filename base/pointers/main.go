package main

import "fmt"

// pointers in Go

func main() {
	basicPointers()
	newAndMake()
	pointerToStruct()
	pointerReceivers()
	sliceAndMapOfPointers()
}

func basicPointers() {
	fmt.Println("-- basic pointers --")
	a := 42
	p := &a
	fmt.Printf("a=%d, *p=%d, p=%p\n", a, *p, p)

	*p = 100
	fmt.Printf("after *p=100 -> a=%d\n", a)

	q := maybeInt(false)
	if q == nil {
		fmt.Println("q is nil")
	}

	increment(&a)
	fmt.Printf("after increment -> a=%d\n", a)
}

func increment(x *int) {
	if x == nil {
		return
	}
	*x++
}

func maybeInt(ok bool) *int {
	if !ok {
		return nil
	}
	v := 5
	return &v
}

func newAndMake() {
	fmt.Println("-- new and make --")
	p := new(int)
	*p = 7
	fmt.Printf("value from new: %d (addr %p)\n", *p, p)

	s := make([]*int, 2)
	x1 := 1
	x2 := 2
	s[0] = &x1
	s[1] = &x2
	fmt.Printf("slice of pointers: %v -> values: %d, %d\n", s, *s[0], *s[1])
}

type Person struct {
	Name string
	Age  int
}

func pointerToStruct() {
	fmt.Println("-- pointer to struct --")
	p := &Person{Name: "Alice", Age: 30}
	fmt.Printf("person: %+v, addr: %p\n", *p, p)

	p.Age = 31
	fmt.Printf("after birthday: %+v\n", *p)
}

type Counter struct{ n int }

func (c *Counter) Inc()      { c.n++ }
func (c Counter) Value() int { return c.n }

func pointerReceivers() {
	fmt.Println("-- pointer receivers --")
	var c Counter
	c.Inc()
	(&c).Inc()
	fmt.Printf("counter value: %d\n", c.Value())
}

func sliceAndMapOfPointers() {
	fmt.Println("-- slice and map of pointers --")
	nums := []int{10, 20, 30}
	ptrs := make([]*int, len(nums))
	for i := range nums {
		ptrs[i] = &nums[i]
	}
	for _, pp := range ptrs {
		fmt.Printf("addr=%p val=%d\n", pp, *pp)
	}

	m := map[string]*Person{}
	m["bob"] = &Person{Name: "Bob", Age: 40}
	if p, ok := m["bob"]; ok {
		fmt.Printf("map lookup -> %+v (addr %p)\n", *p, p)
	}
}
