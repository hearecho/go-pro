package main

import "fmt"

type Info struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	fmt.Printf("+--------+\n")
	printCenter("用户ID",4)
}
