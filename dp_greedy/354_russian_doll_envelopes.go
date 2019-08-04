package dp_greedy

/*
You have a number of envelopes with widths and heights given as a pair of integers (w, h). One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

What is the maximum number of envelopes can you Russian doll? (put one inside other)

Note:
Rotation is not allowed.

Example:

Input: [[5,4],[6,4],[6,7],[2,3]]
Output: 3
Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
 */

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n <= 1 {
		return n
	}
	a := sort(envelopes, 0, n-1) // 对a使用lis在j上

	dp := make([]int, n)
	dp[0] = 1
	max := 0
	for i := 1; i < n; i++ {
		maxTilI := 0
		for j := 0; j <= i; j++ {
			if a[j][1] < a[i][1] {
				if dp[j] > maxTilI {
					maxTilI = dp[j]
				}
			}
		}
		dp[i] = maxTilI + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func sort(a [][]int, start, end int) [][]int {
	if start >= end {
		return a
	}
	i := start
	j := end
	pivot := a[i]
	for i < j {
		for i < j && (a[j][0] > pivot[0] || (a[j][0] == pivot[0] && a[j][1] <= pivot[1])) {
			j--
			continue
		}
		a[i] = a[j]
		for i < j && (a[i][0] < pivot[0] || (a[i][0] == pivot[0] && a[i][1] >= pivot[1])) {
			i++
			continue
		}
		a[j] = a[i]
	}
	a[i] = pivot
	a = sort(a, start, i-1)
	a = sort(a, i+1, end)
	return a
}