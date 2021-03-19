package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// This is a class method for our Vertex Struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	v := Vertex{X: 1.4, Y: 2.6}
	fmt.Printf("abs: %g\n", v.Abs())
}
