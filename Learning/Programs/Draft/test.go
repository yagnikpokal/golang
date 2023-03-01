/*
input = {{1,3},{1,4},{6,8}},{9,10}   output = {{1,4},{6,8},{9,10}} golang program
*/

package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		if intervals[i][0] <= lastMerged[1] {
			lastMerged[1] = max(intervals[i][1], lastMerged[1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {1, 4}, {6, 8}, {9, 10}}
	merged := mergeIntervals(intervals)
	fmt.Println(merged)
}
