package main

func twoSum(nums []int, target int) []int {
	possibleSolutions := make(map[int]int)
	var sol []int
	for i, v := range nums {

		index, ok := possibleSolutions[v]
		if ok && possibleSolutions[v] != i {
			sol = append(sol, index)
			sol = append(sol, i)
			break
		}

		possibleSolutions[target-v] = i
	}

	return sol
}
