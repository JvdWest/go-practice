package main

import (
	"regexp"
	"strings"
)

func returnMatchingRegexString(s string, regExp string) string {
	reg := regexp.MustCompile(regExp)
	matches := reg.FindAllString(s, -1)
	return strings.Join(matches, "")
}

func checkValidity(password string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if reg.MatchString(password) == false {
		return false
	}
	letters := returnMatchingRegexString(password, "[a-zA-Z]+")
	if len(letters)%2 != 0 {
		return false
	}
	digits := returnMatchingRegexString(password, `\d+`)
	if len(digits)%2 != 1 {
		return false
	}
	return true
}

func Solution8(S string) int {
	// Implement your solution here
	passwords := strings.Split(S, " ")
	longestLength := 0
	for _, password := range passwords {
		if checkValidity(password) {
			length := len(password)
			if length > longestLength {
				longestLength = length
			}
		}
	}
	if longestLength == 0 {
		return -1
	} else {
		return longestLength
	}
}
