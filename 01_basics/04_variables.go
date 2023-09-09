package main

import "fmt"

// var declares variables and can be declared in the package or function level
var c, python bool

func main() {
  fmt.Println(c, python)

  // var may include initializer values
  var i, j int = 1, 2
  fmt.Println(i, j)

  // type can be ommited with initializers
  // type will be infered based on value
  var a, b = "'well hello there'", true
  fmt.Println(a, b)

   // factored declartations can be made like facotred imports
  var (
    factoredString string = "hello"
    factoredInt int = 18
  )
  fmt.Println(factoredString, factoredInt)

  // ONLY IN FUNCTION!!! initialized typless declaration can be shorthand
  // at the outer package level 'var', 'func', 'const', 'import' etc keywords are required
  // type is infered from value
  c, d, e := "short", "hand", 8
  fmt.Println(c, d, e)

  // variables with no explicit value initialize at 0, false, or ""
  var number int
  var word string
  var isTrue bool
  fmt.Println(number, word, isTrue)

  // 'const' functions like 'var', but once declared become immutable
  // const can be a character, string, bool, or numeric value
  const unstoppable string = "immovable"
  fmt.Println(unstoppable)

  // numeric constants are extremely high precision
  const bigValue = 1 << 50 //1 bit shifted left 100 times
  fmt.Println(bigValue)
}
