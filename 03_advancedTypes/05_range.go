package main

import "fmt"

var pow = [] int{1, 2, 4, 8, 16, 32, 64, 128, 256}

func main() {
  // 'range' is a for loop that iterates over a slice or map
  // range returns an 'index' and a 'value' pair declared in for loop init
  for i, v := range pow {
    fmt.Println("index:", i, "value:", v)
  }

  // indes, value, or both may be skipped with 'underscore'
  for _, v := range pow {
    fmt.Println("value:", v)
  }
  for i, _ := range pow {
    fmt.Println("index:", i)
  } 
}
