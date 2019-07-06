package string
/*

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
 */

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	dp := [1001][1001]int{{}}

	max := 1
	start := 0
	end := 0
	for i:=0; i < n; i ++ {
		dp[i][i] = 1
		if i <= n-2 {
			if s[i+1] == s[i] {
				dp[i][i+1] = 1
				max = 2
				start = i
				end = i+1
			}
		}
	}
	//fmt.Println(max, start ,end)

	for i:=n-3; i>=0; i-- {
		for j:=i+2; j< n;j++ {
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				dp[i][j]=1
			} else {
				dp[i][j]=0
			}
			if dp[i][j] == 1 {
				if j-i+1 > max {
					max = j-i+1
					start=i
					end=j
				}
			}
		}
	}
	return s[start:end+1]
}