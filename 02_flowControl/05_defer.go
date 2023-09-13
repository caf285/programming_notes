package main

import "fmt"

func main() {
  // defer executes after surrounding function returns
  // defered arguments are evaluated immediately
  defer fmt.Println("world")
	fmt.Println("hello")

  // defers stack and execute in a first-in-last-out order, like a stack
  // defers can be called after previous statements and will execute on container function end
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
}
