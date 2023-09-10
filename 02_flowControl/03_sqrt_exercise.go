package main

import (
  "fmt"
  "math"
)

// exercise writes a custom square root function
func Sqrt(x float64) (float64) {

  // guess a close arbitrary number
  var z float64 = 1

  // the exercise is unclear about how to terminate with the correct answer
  // loop reduces until lower precision floats are equel
  for float32(z) != float32(z - (z * z - x) / (2 * z)) {
    z -= (z * z - x) / (2 * z)
  }
  return z - (z * z - x) / (2 * z)
}

func main() {

  // compare custom function to built in function
  fmt.Println("my sqrt", Sqrt(2))
  fmt.Println("go sqrt", math.Sqrt(2))
}
