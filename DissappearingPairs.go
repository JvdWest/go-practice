package main

import (
	"regexp"
)

func Solution13(S string) string {

	for {
		tempS := S
		reg := regexp.MustCompile("(AA)+|(BB)+|(CC)+")
		S = reg.ReplaceAllString(S, "")
		if len(tempS) == len(S) {
			return S
		}
	}
	// Implement your solution here
}
