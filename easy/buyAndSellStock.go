package main

func maxProfit(prices []int) int {
	min := prices[0]
	current := 0

	for _, v := range prices {
		if min > v {
			min = v
		}

		if v-min > current {
			current = v - min
		}
	}

	return current
}
