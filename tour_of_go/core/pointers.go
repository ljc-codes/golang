package main

import "fmt"

func main() {

	i, j := 42, 80

	p := &i
	fmt.Printf("value of pointer %d\n", *p)
	// set i to a new value
	*p = *p + 1
	fmt.Printf("new value of pointer %d\n", *p)

	p = &j
	*p = *p / 8
	fmt.Printf("changed j value to: %d\n", *p)

}
