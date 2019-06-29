package sum

/*
Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2


 */
// brute force
// func subarraySum(nums []int, k int) int {
//     n := len(nums)
//     res := 0
//     for i:= 0; i < n; i++ {
//         total := nums[i]
//         if total == k {
//             res++
//         }
//         for j:= i+1; j < n; j++ {
//             total += nums[j]
//             if total == k {
//                 res++
//                 //break
//             }
//         }
//     }
//     return res
// }

func subarraySum(nums []int, k int) int {
	n := len(nums)
	res := 0
	sum := 0
	myMap := map[int]int{} // myMap to store the sum and number's of occurenence of the sum
	for i:= 0; i < n; i++ {
		sum += nums[i]
		if sum == k { // 第一次等于k的情况，加上
			res++
		}
		if val, ok := myMap[sum-k]; ok {
			res += val // 之前出现几次，说明中间段相加就等于k的次数有几次，所以都加上
		}
		myMap[sum] = myMap[sum] + 1 // 累加
	}
	return res
}

// 注意，即使当前已经等于，还要继续往下遍历
// 从头遍历到尾 bruteforce, n2, 肯定是work的

// 遍历一次，用一个map保存到目前为止的sum次数，因为比如前一段相加的sum是a, 后面一段相加的sum是b
// 假如b-a = k，那么意味着a到b这一段只有满足条件的情况出现，那么count就要加上前面一段sum出现a的次数
