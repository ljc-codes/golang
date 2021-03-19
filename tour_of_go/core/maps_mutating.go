package main

import "fmt"

func main() {

	m := make(map[string]int)

	// maps "answer" to 42
	m["answer"] = 42
	fmt.Println(m["answer"])

	// maps "answer" to 32
	m["answer"] = 32
	fmt.Println(m["answer"])

	// delete "answer"
	delete(m, "answer")
	fmt.Println("The value:", m["Answer"])

	// checks value of "Answer", v should be the empty
	// value of int. ok will also give us if it's present
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
