package main

import "fmt"

func changeAge(z *int) int {
  *z = 24
  return *z
}

func main() {
  age := 42
  fmt.Println(changeAge(&age))
}