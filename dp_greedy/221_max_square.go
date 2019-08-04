package dp_greedy

/*
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4
 */

func maximalSquare(matrix [][]byte) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	l := len(matrix[0])
	if l == 0 {
		return 0
	}
	dp := make([][]int, h+1)
	for i := range dp {
		dp[i] = make([]int, l+1)
	}
	//fmt.Println(dp)
	maxLength := 0
	for i := 1; i < h+1; i++ {
		for j := 1; j < l+1; j++ {
			if matrix[i-1][j-1] == '1' {
				minLength := min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				dp[i][j] = minLength + 1
				if maxLength < dp[i][j] {
					maxLength = dp[i][j]
				}
			}
		}
	}
	//fmt.Println(dp)
	return maxLength * maxLength
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}