package main

import "fmt"

func main() {

	var pow = []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	for i, v := range pow {
		fmt.Printf("2^%d=%d\n", i, v)
	}

}
