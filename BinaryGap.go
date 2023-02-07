package main

import (
	"fmt"
	"sort"
	"strings"
)

func binaryGap(N int) int {
	binaryString := strings.Trim(fmt.Sprintf("%b", N), "0")
	zeroSplit := strings.Split(binaryString, "1")
	sort.Slice(zeroSplit, func(i int, j int) bool {
		return len(zeroSplit[i]) > len(zeroSplit[j])
	})
	return len(zeroSplit[0])
}
