/*
Given an array of intervals where
intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the
non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]

Output: [[1,6],[8,10],[15,18]]

Explanation: Since intervals [1,3] and [2,6] overlap,
merge them into [1,6].

*/

package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) []Interval {

	var mergedIntervals []Interval
	mergedIntervals = append(mergedIntervals, intervals[0])

	// Range
	for i := 1; i < len(intervals); i++ {
		previous := mergedIntervals[len(mergedIntervals)-1]
		if intervals[i].Start <= previous.End {

			if intervals[i].End > previous.End {
				previous.End = intervals[i].End
			}
			mergedIntervals[len(mergedIntervals)-1] = previous
		} else {
			// Non-overlaping interval, add it to the result
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}

	return mergedIntervals
}

func main() {
	intervals := []Interval{
		{Start: 1, End: 3},
		{Start: 2, End: 6},
		{Start: 8, End: 10},
		{Start: 15, End: 18},
	}
	mergedIntervals := mergeIntervals(intervals)

	for _, interval := range mergedIntervals {
		fmt.Printf("(%d %d)", interval.Start, interval.End)
	}
}
