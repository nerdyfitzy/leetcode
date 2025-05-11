package main

func majorityElement(nums []int) int {
	majority := len(nums) / 2

	intMap := make(map[int]int)

	for _, v := range nums {
		intMap[v]++
		if intMap[v] > majority {
			return v
		}
	}

	//to make compiler happy
	return 0
}
