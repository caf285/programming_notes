package main

import (
  "fmt"
  "math"
)

// function that takes a 'function' as a 'value', then executes that function with float64s 3 and 4
// 'fn' argument is a function that acceps 2 floats and returns a float
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// functions that return functions may also be used as closures
// closures maintain internal variables that can be manipulated with outside calls and arguments
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

// fibonacci closure
func fibonacci() func() int {
  x := 0
  y := 1
  return func() int {
    out := x
    x = y
    y = out + x
    return out
  }
}

func main() {
  // function saves to value 'hypot'
  hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

  // execute hypot with 5 and 12
	fmt.Println(hypot(5, 12))

  // execute hypot with compute values 3 and 4
	fmt.Println(compute(hypot))

  // execute math.Pow with compute values 3 and 4
	fmt.Println(compute(math.Pow))

  // closure function returned and manipulated via calls to itself
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }

  // fibonacci closure
  newFibonacci := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(newFibonacci())
  }
}
