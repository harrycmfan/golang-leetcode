package string

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	res := ""
	for {
		var add byte
		for j := 0; j < len(strs); j++ {
			//fmt.Println(strs, j)
			if len(strs[j]) == i {
				return res
			}
			if j == 0 {
				add = strs[j][i]
			} else {
				if add != strs[j][i] {
					return res
				}
			}
		}
		res += string(add)
		i++
	}
	return res
}

// Vertical scanning