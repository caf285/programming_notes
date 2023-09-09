package main

import "fmt"

// function declarations include parameters, followed by return type
func add(x int, y int) int {
  return x + y
}

// common parameter types may be combined
func addString(x, y int, z string) int {
  return x + y + len(z)
}

// functions may return multiple results
func swap(x, y string) (string, string) {
  return y, x
}

// return values may be named in the signature, then immediately referenced in the function
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return x, y
}

func main() {
  fmt.Println(add(1, 2))

  fmt.Println(addString(1, 2, "asd"))

  a, b := swap("hello", "world")
  fmt.Println(a, b)

  fmt.Println(split(17))
}
