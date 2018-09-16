package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{12, 43, 100, 96, 23, 43, 11, 22}
	fmt.Println("Before Sorting", arr)
	fmt.Println("After Sorting", sort(arr))
}
func sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := rand.Int() % len(a)

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sort(a[:left])
	sort(a[left+1:])

	return a
}
