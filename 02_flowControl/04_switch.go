package main

import (
  "fmt"
  "time"
)

func main() {
  // go 'switch' does not require break statements
  weekday := time.Now().Weekday()
  switch time.Saturday {
    case weekday + 0:
      fmt.Println("today")
    case weekday + 1:
      fmt.Println("tomorrow")
    default:
      fmt.Println("too far away")
  }

  // conditionals can be in cases
  // case without conditional is true
  today := time.Now().Day()
  switch {
    case today < 5:
      fmt.Println(today, "today is less than 5")
    case today < 15:
      fmt.Println(today, "today is less than 15")
    default:
      fmt.Println(today, "is greater or equal to 15")
  }

  // fallthrough allows next case to be evaluated
  switch {
    case today < 5:
      fmt.Println(today, "today is less than 5")
      fallthrough
    case today < 15:
      fmt.Println(today, "today is less than 15")
      fallthrough
    default:
      fmt.Println(today, "is greater than 15? it might be and it might not ...")
  }
}
