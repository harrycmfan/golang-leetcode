package golang_leetcode

/*
Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.
Note: Length of the array will not exceed 10,000.
 */

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	max := 1
	left := 0
	for right := 1; right < n; right++ {
		if nums[right] > nums[right-1] {
			continue
		}
		l := right - left
		if l > max {
			max = l
		}
		left = right
	}
	lastLenght := n - left // 最后这个容易遗忘
	if lastLenght > max {
		max = lastLenght
	}
	return max
}