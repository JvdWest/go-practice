package main

import "fmt"

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

// Not working

func Solution2(N int) int {
	binaryNumber := fmt.Sprintf("%b", N)
	//P := string(binaryNumber[0])
	//Q := "0"
	//for i, _ := range binaryNumber[1:] {
	//	fmt.Println(i)
	//	if P[i-1] == binaryNumber[i] && binaryNumber[i] == '1' {
	//		Q = replaceAtIndex(Q, rune(binaryNumber[i]), i)
	//		P = replaceAtIndex(P, '0', i)
	//	} else { // if Q[i-1] == binaryNumber[i] && binaryNumber[i] == '1'
	//		P = replaceAtIndex(P, rune(binaryNumber[i]), i)
	//		Q = replaceAtIndex(Q, '0', i)
	//	}
	//}

	var P int32
	var Q int32

	for i := 0; i < 32; i++ {
		// Check least significant bit
		bit := int32(N & 0x80000000)
		if P&1 == bit && bit == 1 {
			Q = Q<<1 | 1
			P = P << 1
		} else if Q&1 == bit && bit == 1 {
			Q = Q << 1
			P = P<<1 | 1
		} else {
			P = P<<1 | bit
			Q = Q << 1
		}
		N = N << 1
		fmt.Printf("%b\n", N)
	}

	fmt.Println(binaryNumber)
	fmt.Println(P)
	fmt.Println(Q)
	return 0
}
