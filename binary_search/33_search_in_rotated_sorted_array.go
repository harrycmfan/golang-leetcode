package binary_search

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
 */

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}
	low := 0
	high := n-1
	for low <= high {
		mid := (low + high)/2
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] { // 注意，这里并不能说明low到mid是递增的，但是可以判断target在不在范围内
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1 // 易错点，因为mid肯定不相等了，所以high和low重新赋值的时候需要更近一位
			} else {
				low = mid + 1
			}
		} else {
			if target <= nums[high] && target > nums[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		//fmt.Println(low, high)
	}
	return -1
}
