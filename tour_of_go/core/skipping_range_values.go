package main

import "fmt"

func main() {
	pow := make([]int, 5)
	for i := range pow {
		pow[i] = 1 << uint(i) // ie 2 ^ i
	}
	// Here you see we can use '_' to skip the index
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
