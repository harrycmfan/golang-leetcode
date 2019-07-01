package sum
/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 */
func twoSumInputArrayIsSorted(numbers []int, target int) []int {
	start := 0
	end := len(numbers)-1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start, end}
		}
		if sum > target {
			end--
		}
		if sum < target {
			start++
		}
	}
	return []int{}
}


// func twoSum(nums []int, target int) []int {
//     myMap := map[int]int{}
//     for i, val := range nums {
//         index, ok := myMap[val]
//         if ok {
//             return []int{index+1, i+1}
//         }
//         myMap[target-val] = i
//     }
//     return []int{}
// }

// 这个是之前用map的方法，这个需要O(n)的space，既然array已经sort了，那么可以用双指针来减少space