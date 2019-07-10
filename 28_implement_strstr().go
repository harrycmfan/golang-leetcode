package golang_leetcode

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().


*/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i:=0; i < len(haystack)-len(needle)+1; i++ {
		//fmt.Println(haystack[i:i+len(needle)])
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}

// 注意判断长度来达到最优解
