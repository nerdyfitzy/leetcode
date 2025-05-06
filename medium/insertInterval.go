package main

import (
	"sort"
)

func sortIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	return intervals
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var finalArray [][]int
	//add new one, then sort it. now all we have to do is merge if they conflict
	intervals = append(intervals, newInterval)
	intervals = sortIntervals(intervals)

	//the values we are gonna keep track of to add to our new array
	low := intervals[0][0]
	high := intervals[0][1]
	//we already got first val so just starting on the 2nd element
	for _, v := range intervals[1:] {
		curLow, curHigh := v[0], v[1]

		//if current one's low is before the past one's high then we have an overlap. that means we 100%
		//dont want to add it to the new array we are making yet. otherwise theres no overlap and we can add
		if curLow <= high {
			//if the new one is higher, we will use that one because it full encapsulates the range
			if curHigh > high {
				high = curHigh
			}
		} else {
			//no overlap == add it to the new array
			finalArray = append(finalArray, []int{low, high})
			low, high = curLow, curHigh
		}
	}

	//add the last one
	finalArray = append(finalArray, []int{low, high})

	return finalArray
}
