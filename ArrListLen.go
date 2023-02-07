package main

import "fmt"

func Solution3(A []int) int {
	var i int
	var count int
	temp := 0b11
	fmt.Println(temp)
	for {
		if A[i] == -1 {
			return count + 1
		}
		i = A[i]
		count++
	}
}
