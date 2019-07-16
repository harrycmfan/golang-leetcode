package pure_math

import "fmt"

/*
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
 */

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	nums := []string{}
	for i:= 1; i<=n; i++ {
		nums = append(nums, fmt.Sprintf("%d", i))
	}
	result := ""
	possibilityArray := make([]int, n-1)
	possibilityArray[len(possibilityArray)-1] = 1
	for i:=len(possibilityArray)-2; i >= 0; i-- {
		possibilityArray[i] = possibilityArray[i+1] * (n-i-1)
	}
	//fmt.Println(possibilityArray)
	k = k-1
	for i:=0; i < len(possibilityArray); i++ {
		rank := k / possibilityArray[i]
		selected := nums[rank]
		result += selected
		nums = append(nums[:rank], nums[rank+1:]...)
		k = k % possibilityArray[i]
		//fmt.Println(k, nums)
	}
	return result + nums[0]
}

// 这道题是纯math
// 就是通过不断把k往下找当前可能性的值在哪个group，然后取出当前的num
// 举例，n=4, 随便选一个数字开头，那么都有3*2*1 = 6 种可能
// 那么k/6 = 就能选出需要哪一个 (这里有一个小trick) 假如k=1~6那么说明起始的是1对吧
// 1在我们剩余的numberList是在index 0,所以我们要k=k-1 那么k=0~5的case k/6就能取到index是0的那个数字了
// 每次遍历取完之后，看看k还剩下多少，比如k=9,我们知道第一个数字是2了，那么剩余数字就是1，3，4在数组里面
// k%6=2(注意k减去了1)，那么2/2(剩下的可能性就是2) = 1，那么取index 1的数字就是3
// 继续2%2=0, 那么0/1=0，取剩下的1，4中第一个
// 结束的时候把最后一个append上去
