package main

import "fmt"

func main() {

  // 'if statements' require an expression followed by a curly brace code block
  if (true) {
    fmt.Println("this is true")
  }

  // if conditions do not require parenthesis
  if 2 < 4 {
    fmt.Println("also true")
  }

  // if statements may include 'short statement'
  // short statement variables only exists in if statement
  if v := "new value"; (len(v) > 5) {
    fmt.Println(v)
  }

  // short statement variables persist through else
  if v := 10; (v < 5) {
    fmt.Println(v, "this will never run")
  } else {
    fmt.Println(v, "... but this will")
  }

  // else may also include if
  if (false) {
    fmt.Println("this will never run")
  } else if (true) {
    fmt.Println("... again, this will run")
  }
}
