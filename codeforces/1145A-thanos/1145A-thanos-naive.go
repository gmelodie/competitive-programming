package main

import (
	"fmt"
)

func IsOrdered(arr []int, arrLen int) bool {
	for i := 0; i < arrLen-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Solve(arr []int, arrLen int) int {
	if IsOrdered(arr, arrLen) {
		return arrLen
	}

	a := Solve(arr[:arrLen/2], arrLen/2)
	b := Solve(arr[arrLen/2:], arrLen/2)

	if a > b {
		return a
	}
	return b
}

func main() {
	var length int
	fmt.Scanf("%d", &length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Println(Solve(arr, length))
}
