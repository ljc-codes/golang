package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	var m = map[string]Vertex{
		"simons_foundation": Vertex{
			Lat: 31.5, Long: 95.2},
		"moma": Vertex{
			Lat: 10.2, Long: 98.1},
	}
	fmt.Println(m)
	fmt.Println(m["simons_foundation"])
	fmt.Println(m["moma"])
}
