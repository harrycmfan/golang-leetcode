package sum

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