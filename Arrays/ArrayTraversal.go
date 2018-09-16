package main

import (
	"fmt"
)

func main() {
	var array [5]int
	for i := 0; i < 5; i++ {
		fmt.Print("Enter Number")
		fmt.Scan(&array[i])
	}
	fmt.Println(array)

}
