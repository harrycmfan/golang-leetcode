package dp_greedy

import "strconv"

/*
Given a non-negative integer, you could swap two digits at most once to get the maximum valued number. Return the maximum valued number you could get.

Example 1:
Input: 2736
Output: 7236
Explanation: Swap the number 2 and the number 7.
Example 2:
Input: 9973
Output: 9973
Explanation: No swap.
Note:
The given number is in the range [0, 108]
 */

func maximumSwap(num int) int {
	if num < 10 {
		return num
	}
	nums := []byte(strconv.Itoa(num))
	allDigits := make([]int, 10)
	for i := 0; i < len(nums); i++ {
		allDigits[int(nums[i]-48)] = i
	}
	//fmt.Println(allDigits)
	for i := 0; i < len(nums); i++ {
		for d := 9; d > int(nums[i]-48); d-- {
			if allDigits[d] > i {
				tmp := nums[i]
				nums[i] = nums[allDigits[d]]
				nums[allDigits[d]] = tmp
				res, _ := strconv.Atoi(string(nums))
				return res
			}
		}
	}
	return num
}
