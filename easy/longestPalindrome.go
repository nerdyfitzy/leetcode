package main

func longestPalindrome(s string) int {
	charMap := make(map[rune]int)
	length := 0

	for _, v := range s {
		charMap[v] += 1
		if charMap[v]%2 == 0 {
			length += 2
		}
	}
	if length != len(s) {
		return length + 1
	} else {
		return length
	}
}
