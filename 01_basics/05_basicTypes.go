package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

func main() {
  // classic 'bool' ... or is it? it can only be one or the other
  var isTrue bool = false
  fmt.Println(isTrue)

  // 'string'
  var word string = "'hello my dudes!!'"
  fmt.Println(word)

  // 'int' holds 32 or 64 bits and is best used for index, length, or capacity
  // 'int8', 'int16', 'int32', 'int64' holds 8, 16, 32, and 64 bits respectively and should be used for data
  // 'byte' is alias for 'int8' and 'rune' is alias for 'int32'
  var index int = -5
  var duration int64 = math.MaxInt64
  fmt.Println(index, duration)

  // 'uint', 'uint8', 'uint16', 'uint32', 'uint64' are unsigned versions of above and must be > 0
  var unsignedIndex uint = 5
  fmt.Println(unsignedIndex)

  // 'float32' and 'float64' for decimal numeric values
  var fraction float32 = 12.5876
  fmt.Println(fraction)

  // 'complex64', 'complex128' complex numbers with imaginary counterparts
  var z complex128 = cmplx.Sqrt(-5 + 12i)
  fmt.Println(z)

  // types may be converted with T(v) where value 'v' converts to type 't'
  var f float32 = -1.2
  var i int = int(f)
  var u uint8 = uint8(f)
  fmt.Println(f, i, u)

  // untyped constants take type needed in context
  const bigValue = 1 << 50 //1 bit shifted left 100 times
  fmt.Println(bigValue, bigValue * 0.1)
}
