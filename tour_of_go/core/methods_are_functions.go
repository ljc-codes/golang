package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Observe that methods are just functions
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func main() {
	v := Vertex{X: 3.4, Y: 7.8}
	fmt.Printf("magnitude: %g\n", Abs(v))
}
