package main

import "fmt"

func climb(n, total int) int {
	fmt.Println(n)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return total + 1
	} else if n == 2 {
		return total + 2
	}
	one := climb(n-1, total)
	two := climb(n-2, total)
	return one + two
}

func climbStairs(n int) int {
	total, last, secondLast := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			total = 1
		} else if i == 2 {
			total = 2
		} else {
			total = last + secondLast
		}

		last = secondLast
		secondLast = total
	}

	return total
}
