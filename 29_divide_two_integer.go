package golang_leetcode

import "math"

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
 */

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == math.MinInt32 {
		return 1
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == math.MinInt32 {
		return 0
	}
	positive := false
	if (divisor < 0 && dividend < 0) ||(divisor > 0 && dividend > 0) {
		positive = true
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}

	times := 0
	left := dividend
	sum := 0
	for left >= divisor {
		curr := divisor
		lastCurr := curr
		for left >= curr {
			lastCurr = curr
			curr = curr<<1
			if times > 0 {
				times = times<<1
			} else {
				times = 1
			}
		}
		sum += times
		left = left - lastCurr
		times = 0
	}
	if !positive{
		sum = -sum
	}
	return sum
}

// without using int64
// 位运算最快