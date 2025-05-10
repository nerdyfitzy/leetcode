package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r, _ := regexp.Compile("[^a-z0-9]")
	s = r.ReplaceAllString(s, "")
	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
