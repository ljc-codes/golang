// You can declare methods with pointer receivers.

// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

// For example, the Scale method here is defined on *Vertex.

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

// Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

// With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

// Key idea here is that often times we want to construct methods to mutate our struct itself
// here in the reciever we pass a pointer to our struct

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Here we see that Abs() does not act change our vertex
// we simply return a scalar
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}

// Let's define another functions, Scale() which will
// scale our Vertex vector by some constant f.
func (v *Vertex) Scale(f float64) {
	v.X, v.Y = f*v.X, f*v.Y
}

func main() {
	v := Vertex{X: 20, Y: 74}
	fmt.Printf("norm before: %g\n", v.Abs())
	v.Scale(4)
	fmt.Printf("norm after: %g\n", v.Abs())
}
