package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charCounter := make(map[rune]int)
	for _, v := range s {
		charCounter[v] += 1
	}
	for _, v := range t {
		charCounter[v] -= 1
	}
	for _, v := range charCounter {
		if v != 0 {
			return false
		}
	}

	return true
}
