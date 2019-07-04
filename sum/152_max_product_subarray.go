package sum

/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 */

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := nums[0] // max means the max value at the current position
	min := max     // min means the min value at the current position
	result := max  // result save the result until to the position
	for i:= 1; i < n; i++ {
		if nums[i] < 0 {
			tmp := max
			max = min
			min = tmp
		} // if nums[i] < 0, then the current min and max will be exchanged.
		if nums[i] > max * nums[i] {
			max = nums[i]
		} else {
			max = max * nums[i]
		}
		if nums[i] < min * nums[i] {
			min = nums[i]
		} else {
			min = min * nums[i]
		} // calculate the new min and max at the position
		if max > result {
			result = max
		} // compare with result
	}
	return result
}

// 半动态规划，因为当前位置的max和min对下一个index有作用，但是总的result其实没有作用，每次比较更新就行