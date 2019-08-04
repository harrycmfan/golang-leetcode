package dp_greedy

/*
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
 */

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n) // dp表示的是包含到i的那个数字的最长increase subsequence，注意不是整个数组
	dp[0] = 1
	maxLength := 1
	for i := 1; i < n; i++ {
		maxLengthTilI := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] > maxLengthTilI {
					maxLengthTilI = dp[j]
				}
			}
		}
		dp[i] = maxLengthTilI + 1
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	return maxLength
}