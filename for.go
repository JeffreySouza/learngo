package main

import "fmt"

// for is the only looping construct in Go

func main() {
  i := 1 // var i int = 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // while true
  for {
    fmt.Println("loop")
    break
  }
}
