package main

import "fmt"

// no ternary if in go

func main() {

  num := 7
  if num%2 == 0 {
    fmt.Printf("num %d is even\n", num)
  } else {
    fmt.Printf("num %d is odd\n", num)
  }

  if num = 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
