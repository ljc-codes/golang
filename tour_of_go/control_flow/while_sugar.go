package main

import "fmt"

func main() {

	// It turns out that for is go's while
	// that's awesome

	sum := 0
	i, limit := 0, 400
	for i < limit {
		sum += i ^ 2
		i += 1
	}
	fmt.Printf("got :%d\n", sum)

}
