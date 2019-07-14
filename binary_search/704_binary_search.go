package binary_search

/*
Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
 */

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	//return binarySearch(nums, 0, lenght-1, target)

	left := 0
	right := length - 1

	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return -1
}

func binarySearch(nums []int, start, end int, target int) int {
	if start > end {
		return -1
	}
	if start == end {
		if target == nums[start] {
			return start
		}
		return -1
	}
	mid := (start + end)/2
	midVal := nums[mid]
	if midVal == target {
		return mid
	}
	if midVal > target {
		return binarySearch(nums, start, mid-1, target)
	}
	return binarySearch(nums, mid+1, end, target)
}

// 双指针和递归