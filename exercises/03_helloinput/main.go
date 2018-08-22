package main

import "fmt"

var name string

func main() {
  fmt.Println(name)
  fmt.Scan(&name)
  mem_address := &name
  fmt.Println(mem_address)
  fmt.Println("Hello " + name)
  fmt.Println("Hello " + *mem_address)
  fmt.Println(&name)
}

