package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter First Number:=")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number:=")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Print(num1, " is Greater than ", num2, " and ", num2, " is Less than ", num1)
	} else if num1 == num2 {
		fmt.Print(num1, " is Equal to ", num2)
	} else if num1 < num2 {
		fmt.Print(num2, " is Greater than ", num1, " and ", num1, " is Less than ", num2)
	}
}
