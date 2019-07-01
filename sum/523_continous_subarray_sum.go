package sum

/*
Given a list of non-negative numbers and a target integer k, write a function to check if the array has a continuous subarray of size at least 2 that sums up to a multiple of k, that is, sums up to n*k where n is also an integer.



Example 1:

Input: [23, 2, 4, 6, 7],  k=6
Output: True
Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
Example 2:

Input: [23, 2, 6, 4, 7],  k=6
Output: True
Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
 */

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	leftMap := map[int]int{}
	sum := nums[0]
	if k != 0 {
		left := nums[0]%k
		leftMap[left] = 0
	}
	for i:=1; i< len(nums); i++ {
		if nums[i] == 0 {
			if nums[i-1] == 0 {
				return true
			}
			continue
		}
		sum += nums[i]
		if k != 0 {
			left := sum%k
			//fmt.Println(leftMap, sum, left)
			if left == 0 { //没有余数，当然ok
				return true
			}
			lastIndex, ok := leftMap[left]
			if ok && lastIndex + 1 != i {
				return true
			}
			leftMap[left] = i
		}
	}
	return false
}

// func checkSubarraySum(nums []int, k int) bool {
//     n := len(nums)
//     if n <= 1 {
//         return false
//     }
//     sumNums := make([]int, n)
//     sumNums[0] = nums[0]
//     sum := nums[0]
//     for i:=1; i< len(nums); i++ {
//         if nums[i] == 0 {
//             if nums[i-1] == 0 {
//                 return true
//             }
//             continue
//         }
//         sum += nums[i]
//         //fmt.Println(sum)
//         for j:=0; j< i-1; j++ { //注意相邻的两个如果相减可以整除，说明只有一个数，所以这里j不能取i-1
//             if k != 0 && (sum - sumNums[j])%k == 0 {
//                 return true
//             }
//             if k == 0 && sum == 0 && sumNums[j] == 0 {
//                 return true
//             }
//         }
//         sumNums[i] = sum
//         if k != 0 && sum%k == 0 {
//             return true
//         }
//     }
//     return false
// }

// bruteforce, 记录一下每个index目前为止的sum，每遍历到某一个index，就看当前sum减去之前的每一次的sum是不是能被k整除，是的话就return true(注意0，和相减不能使最近一个的edge case)，这是n2. (这题的k=0个人认为有点问题，理解就行了)

// 既然在遍历的时候我们保存了每个index目前为止的sum，而遍历到某一个index的时候，即使这个index下的sum并不能呗整除，我们还是会知道余数，如果余数相同，那么之间的差值肯定可以整除，注意edgecase就好了
// 这题其实就是sum array equals k的进阶

