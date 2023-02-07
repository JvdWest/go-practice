package main

import (
	"fmt"
	"math"
	"strings"
)

func supers(N int) int {
	//newString := fmt.Sprintf("%030b\n", 5)
	//fmt.Println(strings.Count(newString, "0"))
	fmt.Printf("%030b\n", 5)
	numberOfZeros := float64(strings.Count(fmt.Sprintf("%030b", N), "0"))
	return int(math.Pow(2, numberOfZeros))
}

func Solution1(A int, B int, C int) int {
	totalCombinations := supers(A) + supers(B) + supers(C)
	totalCombinations -= supers(A | B)
	totalCombinations -= supers(A | C)
	totalCombinations -= supers(B | C)
	totalCombinations += supers(A | B | C)
	return totalCombinations
}
