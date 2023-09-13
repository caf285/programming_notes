package main

import "fmt"

func main() {
  i, j := 50, 200

  // point to i with '&i' and access i from '*p'
  p := &i // p = &i = address of i
  fmt.Println("p:", p, "*p:", *p, "i:", i)
  *p = 25 // *p = i = value of &i
  fmt.Println("p:", p, "*p:", *p, "i:", i)

  // point to j and perform math through pointer
  p = &j // p = &j = address of j
  fmt.Println("p:", p, "*p:", *p, "j:", j)
  *p += *p // *p = j = value of &j
  fmt.Println("p:", p, "*p:", *p, "j:", j)
}
