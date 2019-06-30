package sum

import "sort"

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n <= 3 {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	sort.Ints(nums)
	closestMax := nums[n-1] + nums[n-2] + nums[n-3] - target
	closestMin := target - (nums[0] + nums[1] + nums[2]) // 注意括号
	//fmt.Println(nums, closestMax, closestMin)
	for i:= 0; i < n-2; i++ {
		left := i+1
		right := n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum > target {
				if sum - target < closestMax {
					closestMax = sum - target
				}
				right--
			} else {
				if target - sum < closestMin { // 定义了两个closest 目的就是让他是大于0的整数，所以这里还是直接判断target- sum <
					closestMin = target - sum
				}
				left++
			}
		}
	}
	if closestMax < closestMin {
		return closestMax + target // 注意，题目要求回sum不是closest的值
	}
	return target - closestMin
}