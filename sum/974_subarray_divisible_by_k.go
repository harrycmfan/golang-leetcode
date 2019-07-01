package sum

/*
Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.



Example 1:

Input: A = [4,5,0,-2,-3,1], K = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by K = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]


Note:

1 <= A.length <= 30000
-10000 <= A[i] <= 10000
2 <= K <= 10000
 */

func subarraysDivByK(a []int, k int) int {
	if k == 0 {
		return 0
	}
	if k < 0 {
		k = -k
	}
	n := len(a)
	if n < 1 {
		return 0
	}
	preSumArray := make([]int, n)
	preSumArray[0] = a[0]
	for i:=1; i < n; i++ {
		preSumArray[i] = preSumArray[i-1] + a[i]
	}
	res := 0
	myMap := map[int]int{0:1} // key: left, value: occurrency of left
	//fmt.Println(preSumArray)
	for i := 0; i < n; i++ {
		left := preSumArray[i] % k
		if left < 0 {
			left += k
		}
		if occurency, ok := myMap[left]; ok {
			res += occurency
		}
		myMap[left] = myMap[left] + 1
	}
	//fmt.Println(myMap)
	return res
}

// 借用map的形式类似于subarray sum equals k,但是这里需要想到的是preSum的余数如果相等，意味着两段可以加起来