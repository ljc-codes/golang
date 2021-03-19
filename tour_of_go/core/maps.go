package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)

	v1 := Vertex{Lat: 31.5, Long: 51.9}

	v2 := Vertex{Lat: 91.4, Long: 25.1}

	m["vertex_1"] = v1
	m["vertex_2"] = v2

	fmt.Println(m)

	fmt.Println(m["vertex_1"])
	fmt.Println(m["vertex_2"].Lat)
	fmt.Println(m["vertex_2"].Long)

}
