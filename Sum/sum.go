package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Enter First Number:=")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number:=")
	fmt.Scan(&num2)
	fmt.Print("The Sum of", num1, " and ", num2, " is:= ", num1+num2)
}
