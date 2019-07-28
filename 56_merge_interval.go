package golang_leetcode

/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
 */

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	intervals = sortIntervals(intervals)
	//fmt.Println(intervals)
	result := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end { // attention here to only extend if the following last is bigger
				end = intervals[i][1]
			}
		} else {
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}

func sortIntervals(list [][]int) [][]int {
	return quickSort(list, 0, len(list)-1)
}

func quickSort(list [][]int, left, right int) [][]int { // remember quick sort
	if left  >= right {
		return list
	}
	pivot := list[left]
	i := left
	j := right
	for i < j {
		for i < j && list[j][0] >= pivot[0] {
			j--
		}
		list[i] = list[j]
		for i < j && list[i][0] <= pivot[0] {
			i++
		}
		list[j] = list[i]
	}
	list[i] = pivot
	list = quickSort(list, left, i-1)
	list = quickSort(list, i+1, right)
	return list
}