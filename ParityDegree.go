package main

import (
	"math"
)

func Solution7(N int) int {
	// Implement your solution here
	K := int(math.Ceil(math.Log2(float64(N))))
	for {
		powValue := int(math.Pow(float64(2), float64(K)))
		if K < 0 {
			return 0
		}
		if N%powValue == 0 {
			return K
		}
		K--
	}
}
