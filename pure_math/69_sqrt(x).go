package pure_math

/*

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
 */


func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	if x < 9 {
		return 2
	}
	left := 3
	right := x/3
	for left < right {
		mid := (left + right)/2
		squre := mid * mid
		if squre == x {
			return mid
		} else if squre > x || squre < 0 { // 为什么不用x/mid 与mid比较呢？这样就没有溢出了
			right = mid // 大于 那么right可以是mid-1
		} else {
			if left == mid {
				return left
			} // 如果right = mid+1, left=mid-1那么要注意退出条件
			left = mid // left不可以是mid+1,因为mid+1的话有可能大于x了
		}
		//fmt.Println(left, right)
	}
	return left
}