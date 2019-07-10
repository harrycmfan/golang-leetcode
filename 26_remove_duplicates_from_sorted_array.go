package golang_leetcode

/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
 */

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := len(nums)
	previous := nums[0]
	for i:=1; i< len(nums); i++ {
		if nums[i] == previous {
			count--
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			previous = nums[i]
		}
		//fmt.Println(nums, i, previous)
	}
	return count
}

// 这个并不是最优解哦~，题目要求只要nums的前几位是对的就行了，所以不用concat nums 只需要往前赋值即可