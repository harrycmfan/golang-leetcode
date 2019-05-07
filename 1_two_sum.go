package golang_leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

func twoSum(nums []int, target int) []int {
	myMap := map[int]int{}
	for i, val := range nums {
		index, ok := myMap[val]
		if ok {
			return []int{index, i}
		}
		myMap[target-val] = i
	}
	return []int{}
}

// 傻办法，选第一个elem，遍历后面所有，如果满足return，n2

// hashtable，遍历一次，把所有的elem的差值作为key存到map里面，value是index
// 再遍历一次，看看有没有找到的，返回两个index 2n
// 看代码，根本不用再遍历了，因为之前的已经存过了 所以就是n