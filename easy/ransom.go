package main

func canConstruct(ransomNote string, magazine string) bool {
	counter := make(map[rune]int)

	for _, v := range magazine {
		counter[v] += 1
	}

	for _, v := range ransomNote {
		counter[v] -= 1
	}

	for _, v := range counter {
		if v < 0 {
			return false
		}
	}
	return true
}
