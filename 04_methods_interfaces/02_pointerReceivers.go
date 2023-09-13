package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// extend type Vertex with method
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// manipulate Vertex instance directly by 'receiving pointer'
// no return value necessary when only manipulating pointer values
/*
    (v *Vertex) is a pointer to the instance of Vertex 'pass by reference'
    (v Vertex) is 'pass by value' which means manipulating struct components will have no effect
*/
func (v *Vertex) Scale(f float64) {
  v.X *= f
  v.Y *= f
}

// both Abs and Scale may also be written as functions
// functions are less efficient as values may be copied for every call
func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
  v.X *= f
  v.Y *= f
}

func main() {
  // call methods
  // Scale is called with no change
  v := Vertex{3, 4}
  v.Scale(10) // automatically interpreted as (&v).Scale(10)
  fmt.Println(v.Abs())
  p := &v
  p.Scale(10) // automatically interpreted as (*p.Scale(10))
  fmt.Println(p.Abs())

  // call functions
  // functional scale requires passing address of v as in '&v'
  v = Vertex{3, 4}
  Scale(&v, 10)
  fmt.Println(Abs(v))
}
