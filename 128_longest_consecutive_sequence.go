package golang_leetcode

/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
 */

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	max := 1
	myMap := map[int]int{}
	for _, num := range nums {
		myMap[num] = 1
	}
	for _, num := range nums {
		if _, ok := myMap[num-1]; ok {
			continue
		}
		consecutive := 1
		cur := num
		for {
			if _, ok := myMap[cur+1]; ok {
				cur++
				consecutive++
			} else {
				break
			}
		}
		if consecutive > max {
			max = consecutive
		}
	}

	return max
}

// idea就是把所有的number存到map中，然后遍历，如果不是初始的num那么就往下找，这个地方非常clever