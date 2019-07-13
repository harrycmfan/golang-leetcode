package string

/*
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.



Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False


Note:

The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].
 */

func checkInclusion(s1 string, s2 string) bool {
	indicator := make([]int, 128)
	for _, b := range s1 {
		indicator[b]++
	}

	left := 0
	for right := 0; right < len(s2); right++ {
		b := s2[right]
		if indicator[b] > 0 {
			indicator[b]--
			if right - left + 1 == len(s1) {
				return true
			}
		} else {
			for left <= right {
				leftB := s2[left]
				if leftB == b {
					break
				}
				indicator[leftB]++
				left++
			}
			left++
		}
	}
	return false
}

// 必须是双指针因为当right出现重复的时候，left可以向右移动