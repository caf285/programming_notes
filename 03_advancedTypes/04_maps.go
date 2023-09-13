package main

import "fmt"

// alternatively 'X, Y int'
type Vertex struct {
  X int
  Y int
}

// map is like struct, but keys are required
// map literal is 'map[T1]T2' as 'T1' is used to reference 'T2'
var m1 = map[string]Vertex{
  "Bell Labs": Vertex{
    40, -74,
  },
  "Google": Vertex{
    37, -122,
  },
}

// if top-level type is just the reference type name, it may be omited
var m2 = map[string]Vertex{
  "Bell Labs": { 40, -74 },
  "Google": { 37, -122 },
}

func main() {
  fmt.Println(m1)
  fmt.Println(m2)

  // maps may also be defined with 'make'
  m3 := make(map[string]int) // empty map of m{string: int}

  // maps may be mutated by 'adding/updating' key/values or 'deleting' keys
  m3["one"] = 1
  m3["two"] = 2
  m3["three"] = 3
  fmt.Println(m3)
  delete(m3, "three")
  fmt.Println(m3)

  // map values can be reference with map keys
  fmt.Println(m3["one"], m3["two"])

  // a map can be tested for the inclusion of a key with a two-value assignment
  value, ok := m3["one"] // retrives value
  fmt.Println(value, ok)
  value, ok = m3["three"] // returns '0' and 'false' for non-existant key
  fmt.Println(value, ok)

  // 'underscore' discards returned value and does not need to be declared
  _, ok = m3["three"]
}
