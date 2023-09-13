package main

import (
  "fmt"
  "math"
)

// define custom type
type Vertex struct {
  X, Y float64
}

// define custom 'method' for Vertex called 'Abs'
// Vertex 'v' is the method receiver 
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// a method is just a function with a receiver argument
// here is Abs in function format
func Abs(v Vertex) float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y) 
}

// a method can be used to extend built in types
// !!! cannot use a type declared outside the package including built in types like 'int'
// the type must first be extended such as 'float64' to 'MyFloat'
type MyFloat float64
func (f MyFloat) Abs() float64 {
  if (f < 0) {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  // create 'Vertex' instance and execute method 'Abs'
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  fmt.Println(Abs(v))

  // output method of float64 extansion
  f := MyFloat(-math.Sqrt(2))
  fmt.Println(f.Abs())
}
