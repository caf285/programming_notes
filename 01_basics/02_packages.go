package main

// factored imports combine multiple import calls
import (
  "fmt"
  "math"
)

func main() {
  // exported names begin with capitol letter
  // 'Pi' is an exported function from the 'math' package
  fmt.Println(math.Pi)
}
