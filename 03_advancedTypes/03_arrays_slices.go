package main

import "fmt"

func main() {
  // 'array' is declared '[n]T'; 'n' values of type 'T'
  // array length is required for array and is immutable
  var a [10]int
  fmt.Println(a)

  // array indices can be access via square brackets
  b := [5]string{"world", "hello", "1", "2", "3"}
  b[0] = b[1]
  b[1] = "world"
  fmt.Println(b[0], b[1])

  // copying an array will copy by value to a new memory address
  c := b
  fmt.Println(&b[0], &c[0])  
  c[0] = "goodbye"
  fmt.Println(b, c)

  // 'slice' is a dynamically sized view of an array
  // slice is defined with no 'type length', but low and high 'bounds' of reference array
  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  fmt.Println(s)

  // slices do not store data, but reference an array
  // s references primes[1:4] and s[0] references primes[1]
  s[0] = 50
  fmt.Println(primes, &primes[1], &s[0])

  // slice literal creates array and references array
  array := [3]int{1, 2, 3} // creates array
  slice := []int{1, 2, 3} // creates array, then references array
  fmt.Println(array, slice)

  // structs can be defined per slice on the fly
  structSlice := []struct {
    i int
    b bool
  }{
    {15, true},
    {100, false},
  }
  fmt.Println(structSlice)

  // appending a slice
  // re-slicing beyond the bounds of underlying array doubles new array size as needed
  appSlice := []int{1, 2, 3} // underlying array is {1, 2, 3}
  appSlice = append(appSlice, 4) // slice is now {1, 2, 3, 4}, but new array is {1, 2, 3, 4, 0, 0}

  // slice copy references same array
  // bounds may also be omitted
  oldSlice := []int{1, 2, 3}
  newSlice := oldSlice[:2]
  fmt.Println(&oldSlice[0], &newSlice[0])

  // slice copy references new array after appending outside boudes of original
  newSlice = append(newSlice, 3)
  fmt.Println(&oldSlice[0], &newSlice[0], oldSlice, newSlice) // same address
  newSlice = append(newSlice, 4)
  fmt.Println(&oldSlice[0], &newSlice[0], oldSlice, newSlice) // new array, new address

  // 'length' of a slice is the number of elements for a slice 'len(S)'
  // 'capacity' of a slice is the number of elements in the underlying array 'cap(S)'
  capacitySlice := []int{1, 2, 3, 4, 5}
  fmt.Println(len(capacitySlice), cap(capacitySlice))

  // length and capacity begin counting from 0 index of slice and not array
  capacitySlice = capacitySlice[1:]
  fmt.Println(len(capacitySlice), cap(capacitySlice)) // previous indices of array are lost

  // increasing a slice beyond the underlying array creates a new array with double the capacity
  subSlice := capacitySlice[:1]
  fmt.Println(&subSlice[0], &capacitySlice[0]) // same address
  capacitySlice = append(capacitySlice, 6) // slice needs new larger underlying array
  fmt.Println(&subSlice[0], &capacitySlice[0]) // address is now different
  fmt.Println(len(capacitySlice), cap(capacitySlice)) // new array capacity now doubled

  // 'nil slice' have capacity of 0 and no underlying array
  var nilSlice []string
  fmt.Println(nilSlice, nilSlice == nil) // nil slice has no underlying array
  nilArray := [2]string{"nil", "array"}
  nilSlice = nilArray[2:]
  fmt.Println(nilSlice, nilSlice == nil) // empty slice, but not nil slice, because underlying array
  
  // define underlying array parameters for slice with 'make'
  // make literal is make(T len ...cap)
  mkSlice1 := make([]int, 5) // len of 5 ... with default cap of 5
  mkSlice2 := make([]int, 0, 5) // len of 0, but underlying array cap of 5
  mkSlice2 = mkSlice2[:cap(mkSlice2)] // extend slice to length of underlying array
  mkSlice2 = mkSlice2[1:] // len and cap now 4
  fmt.Println(mkSlice1, mkSlice2)

  // slices of slices or 'embedded slice'
  // inner slices can be of any length
  embSlice := [][]string{
    []string{"hello", "1"},
    []string{"hello", "2", "3"},
  }
  fmt.Println(embSlice)
}
