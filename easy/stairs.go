package main

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
