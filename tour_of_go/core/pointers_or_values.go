// There are two reasons why we'd pass by reference vs values.
//
// 1) The first is so that the method can modify the value that its receiver points to.
//
// 2) The second is to avoid copying the value on each method call. This can be more
//    efficient if the receiver is a large struct, for example.
//
// Generally we ought not to combine pointers and values
//
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X, v.Y = f*v.X, f*v.Y
}

func main() {

	vertex := Vertex{X: 23.1, Y: 68.2}
	fmt.Printf("val: before: %g\n", vertex.Abs())
	vertex.Scale(5)
	fmt.Printf("val: after: %g\n", vertex.Abs())

	// Notice that passing the reference also works
	d := &vertex
	fmt.Printf("ref: before: %g\n", d.Abs())
	vertex.Scale(10)
	fmt.Printf("ref: after: %g\n", d.Abs())

}
