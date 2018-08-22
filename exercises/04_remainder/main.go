package main

import "fmt"

var largeNum int
var smallNum int

func main() {
	fmt.Printf("%v", "Please Enter a Large Number: ")
	fmt.Scan(&largeNum)
	fmt.Printf("%v", "Please Enter a Small Number: ")
	fmt.Scan(&smallNum)
	fmt.Println(largeNum % smallNum)
}
