package sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	firstFour := nums[0] + nums[1] + nums[2] + nums[3]
	if firstFour == target {
		return [][]int{{nums[0],nums[1],nums[2],nums[3]}}
	}
	if firstFour > target {
		return [][]int{}
	}
	result := [][]int{}
	for i:= 0; i < n - 3; i++ {
		for j:=i+1; j < n - 2; j++ {
			left := j+1
			right := n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left+1] == nums[left] {
						left++
					}
					for left < right && nums[right-1] == nums[right] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
			for j < n-2 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i < n-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}


// k sum O(k-1) 没有别的办法，从头遍历，最后用双指针
