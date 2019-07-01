package sum

/*Your are given an array of positive integers nums.

Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.

Example 1:
Input: nums = [10, 5, 2, 6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
Note:

0 < nums.length <= 50000.
0 < nums[i] < 1000.
0 <= k < 10^6.
 */

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	n := len(nums)
	if n == 1 {
		if nums[0] < k {
			return 1
		} else {
			return 0
		}
	}
	left := 0
	product := nums[left]
	res := 0
	if product < k {
		res++
	}
	for right := 1; right < n; right++ {
		product *= nums[right]
		//fmt.Println(product)
		for product >= k && left < n{
			product /= nums[left]
			left++
			//fmt.Println(left)
		}
		if left == n {
			return res
		} else if left > right {
			res++
			right = left
		} else {
			//fmt.Println(left, right)
			res += right-left+1 // 增加所有的cases which include right
		}
		// fmt.Println(res)
	}
	return res
}

// sliding window
