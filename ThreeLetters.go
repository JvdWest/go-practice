package main

import "strings"

func completeUniformSequence(count int, lastChar string) string {
	if count == 0 {
		return ""
	}
	var result string
	var addedString string
	if lastChar == "a" {
		addedString = "ba"
	} else {
		addedString = "ab"
	}
	for i := count; i > 0; i-- {
		result += addedString
	}
	return result
}

func Solution6(A int, B int) string {
	if A == B {
		return completeUniformSequence(A, "a")
	}
	var answer string
	for {
		if A > B {
			answer += "aa" + "b"
			A -= 2
			B -= 1
			if A == B {
				return answer + completeUniformSequence(A, "b")
			}
		} else {
			answer += "bb" + "a"
			A -= 1
			B -= 2
			if A == B {
				return answer + completeUniformSequence(A, "a")
			}
		}
		if A == 0 {
			return answer + strings.Repeat("b", B)
		}
		if B == 0 {
			return answer + strings.Repeat("a", A)
		}
	}
}
