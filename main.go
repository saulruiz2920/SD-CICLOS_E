package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  var e float64 = 1
  var fact int64 = 1
  for i := 1; i <= n ; i++ {
    fact *= int64(i)
    e += float64(float64(1) / float64(fact))
  }
  fmt.Println(e)
}