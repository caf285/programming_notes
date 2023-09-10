package main

import "fmt"

func main() {
  // for loops contain 'init', 'condition', and 'post' statements
  // like 'if short statements', for loop parameters cannot use parenthesis
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  // init and post statements are optional
  sum := 1
  for ; sum < 100; {
    sum++
  }
  fmt.Println("sum is now", sum)

  // 'while loops' can be utilized by not including semicolons around the conditional
  sum = 0
  for sum < 100 {
    sum++
  }
  fmt.Println("sum is now", sum)

  // no conditionals will allow a for loop to loop forever
  for {
    fmt.Println("I will loop forever ... unless someone included a handy loop break")
    break
  }
  fmt.Println("... they did :D")

  // 'continue' and 'break' may be used to manipulate a for loop
  for i:= 0; i < 10; i++ {
    if (i < 3) {
      fmt.Println(i, "is less than 3")
    }
    if (i >= 5) {
      fmt.Println("i is", i, "... lets break early")
      break
    }
  }
}
