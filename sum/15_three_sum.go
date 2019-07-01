package sum

import "sort"

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}
	if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
		return [][]int{{0,0,0}}
	}

	result := [][]int{}
	for i:=0; i<len(nums);i++ {
		if nums[i] > 0 {
			break
		}
		left := i+1
		right := len(nums)-1
		for ;left < right; {
			if nums[i] + nums[left] + nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < len(nums)-1 && nums[left] == nums[left+1] {
					left++
				}
				for right > 0 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

// 最笨的办法，3次遍历，n3 (这里其实还不是很好处理，因为还要去重)
// 首先第一印象必须要sort，因为只有排序好了，才方便后面遍历来取和
// 排序后，对于第一个数，后面通过left，right 2 pointer来相加取和，注意相同的数值要跳过，避免重复
// 同理找到一波组合之后，需要判断当前后面的数是不是跟现在一样，是就要跳过
// 这个办法一定是比用set要好的，因为用set去重其实去重本身也有complexity

