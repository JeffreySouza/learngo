package main

import "fmt"

func main() {

  var a string = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c) // adds spaces between numbers, nice

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e) // uninitialized vars are zero-valued, int is 0

  f := "short" // short hand for var f string = "short"
  fmt.Println(f)
}
