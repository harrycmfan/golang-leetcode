package golang_leetcode

/**
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.


*/

// func findKthLargest(nums []int, k int) int {
//     sort.Ints(nums)
//     n := len(nums)
//     return nums[n-k]
// }

func findKthLargest(nums []int, k int) int {
	// if len(nums) == 1 {
	//     return nums[0]
	// }
	left := 0
	right := len(nums)-1
	for left < right {
		//fmt.Println(left, right)
		pivotPos := partition(nums, left, right)
		//fmt.Println(nums)
		rank := len(nums) - pivotPos // 这里可以得到pivot是第rank大的
		if rank == k {
			break
			//return nums[pivotPos]
		} else if rank > k { // 说明pivot比要找的小了，需要找更大的
			left = pivotPos + 1 // left在pivotPos后面的都是比pivot的值大的
		} else { // 大了，需要找更小的
			right = pivotPos - 1
		}
	}
	// 如果k一定是valid的，那么出来的时候k这个位置已经是partition过了 排好了
	//fmt.Println(nums)
	return nums[len(nums)-k]
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	l := start
	r := end
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	// 现在l就表示原本的nums[start]假如排序之后应该处于的位置，
	// 那么右边就有len(nums)-l-1个比他大的,他就是第len(nums)-l大的元素
	return l // return pivotPos
}

// quick select的思想
