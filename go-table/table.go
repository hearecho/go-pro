package main

import (
	"fmt"
	"unicode/utf8"
)

func CreateTable(headers []string)  {
	printCenter("用户ID",7)
}

func printCenter(s string,n int) {
	length := utf8.RuneCountInString(s)
	fmt.Println(length)
	fmt.Printf("|%*s%*s",n+length/2,s,n-length/2," ")
}
