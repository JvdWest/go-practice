package main

import (
	"fmt"
	"sort"
)

func minMax(A []int) (minValue, maxValue int) {
	minValue = A[0]
	for _, val := range A {
		if maxValue < val {
			maxValue = val
		}
		if minValue > val {
			minValue = val
		}
	}
	return
}

func Solution(K int, A []int) (result int) {
	sort.Ints(A)
	for Q, _ := range A {
		for P := 0; P <= Q; P++ {
			min, max := A[P], A[Q]
			if max-min <= K {
				fmt.Println(P, Q)
				result++
			}
		}
	}
	return
}
