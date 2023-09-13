package main

import "fmt"

// new 'struct' makes custom types
// struct components define collection of variables and types
type Vertex struct {
  X int
  Y int
}

func main() {

  // struct fields are accessible with 'dot' operator
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v, v.X)

  // struct fields can be accessed via pointer
  // pointer shorthand allows access via 'p.X' instead of '(*p).X'
  p := &v
  p.X = 1e4
  fmt.Println(v)

  // struct literals allow listing any field in any order
  var (
    v1 = Vertex{1, 2} // has type Vertex
	  v2 = Vertex{X: 1} // Y:0 is implicit
  	v3 = Vertex{}     // X:0 and Y:0
  	p1= &Vertex{1, 2} // has type *Vertex
  )
  fmt.Println(v1, v2, v3, p1)
  
}
