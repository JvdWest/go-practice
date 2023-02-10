package main

import (
	"fmt"
	"math"
)

func findLocalMaximumIndexes(A []int) (localMaximumsIndices []int) {
	if A[0] > A[1] {
		localMaximumsIndices = append(localMaximumsIndices, 0)
	}
	if A[len(A)-1] > A[len(A)-2] {
		localMaximumsIndices = append(localMaximumsIndices, len(A)-1)
	}
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			localMaximumsIndices = append(localMaximumsIndices, i)
		}
	}
	return
}

type pillarObj struct {
	height int
	index  int
}

func findMaximumIndexes(A []int) (maximums []int) {

	var leftPillar pillarObj
	var rightPillar pillarObj
	var index int
	// Find first pillar
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			leftPillar.height = A[i]
			leftPillar.index = i
			rightPillar.height = A[i+1]
			rightPillar.index = i + 1
			index = i + 2
			break
		}
	}

	for {
		// Stop condition
		if index >= len(A) {
			break
		}
		if A[index] > rightPillar.height {
			rightPillar.index = index
			rightPillar.height = A[index]
		}
		if rightPillar.height >= leftPillar.height {
			maximums = append(maximums, leftPillar.index)
			leftPillar = rightPillar
			rightPillar.index = -1
			rightPillar.height = -1
		}
		index++
	}

	maximums = append(maximums, leftPillar.index)
	maximums = append(maximums, rightPillar.index)

	leftPillar = rightPillar
	rightPillar.index = 0
	rightPillar.height = 0
	// Not the last element, look for dams in the last section
	if leftPillar.index != len(A)-1 {
		for i := leftPillar.index; i < len(A); i++ {
			if A[i] > leftPillar.height {
				rightPillar.index = index
				rightPillar.height = A[index]
			}
		}
	}

	//for _, i := range maximumIndices1 {
	//
	//}

	//var localMaximumsIndices []int
	//if A[0] > A[1] {
	//	localMaximumsIndices = append(localMaximumsIndices, 0)
	//}
	//if A[len(A)-1] > A[len(A)-2] {
	//	localMaximumsIndices = append(localMaximumsIndices, len(A)-1)
	//}
	//for i := 1; i < len(A)-1; i++ {
	//	if A[i] > A[i-1] && A[i] > A[i+1] {
	//		localMaximumsIndices = append(localMaximumsIndices, i)
	//	}
	//}
	//
	//// Remove localMaximumsIndices with
	////maximums = append(maximums, A[0])
	//maximums = append(maximums, localMaximumsIndices[0])
	//var temp []int
	//var maxPoint = A[localMaximumsIndices[0]]
	//for i := 1; i < len(localMaximumsIndices); i++ {
	//	if maxPoint < A[localMaximumsIndices[i]] {
	//		temp = []int{} // Clear
	//		maxPoint = A[localMaximumsIndices[i]]
	//		maximums = append(maximums, localMaximumsIndices[i])
	//	} else {
	//		temp = append(temp, localMaximumsIndices[i])
	//	}
	//}
	//for _, val := range temp {
	//	maximums = append(maximums, val)
	//}
	return
}

func Solution9(A []int) int {
	maximumIndices := findMaximumIndexes(A)
	fmt.Println(maximumIndices)
	var maximumDepth int
	for i := 0; i < len(maximumIndices)-1; i++ {
		leftHeight := A[maximumIndices[i]]
		rightHeight := A[maximumIndices[i+1]]
		hillHeight := int(math.Min(float64(leftHeight), float64(rightHeight)))
		for k := maximumIndices[i] + 1; k < maximumIndices[i+1]; k++ {
			if hillHeight-A[k] > maximumDepth {
				maximumDepth = hillHeight - A[k]
			}
		}
	}
	// Implement your solution here
	return maximumDepth
}
