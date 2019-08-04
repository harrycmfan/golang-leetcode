package dp_greedy

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0] // 注意dp是连续到最后一个index的最大subarray
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			if dp[i-1] <= 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = dp[i-1] + nums[i]
			}
		} else {
			if dp[i-1] + nums[i] >= 0 {
				dp[i] = dp[i-1] + nums[i]
			} else {
				dp[i] = nums[i]
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	//fmt.Println(dp)
	return max
}
