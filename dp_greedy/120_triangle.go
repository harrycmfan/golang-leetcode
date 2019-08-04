package dp_greedy

/*
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

 */

func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	if h == 0 {
		return 0
	}
	if h == 1 {
		return triangle[0][0]
	}
	dp := make([]int, h)
	for i := 0; i < h; i++ {
		dp[i] = triangle[h-1][i]
	}
	for i := h-2; i >= 0; i-- {
		//fmt.Println(111, dp)
		for j:= 0; j <= i; j++ {
			tmp := dp[j+1]
			if dp[j] < tmp {
				tmp = dp[j]
			}
			dp[j] = triangle[i][j] + tmp
			//fmt.Println(222, dp)
		}
	}
	return dp[0]
}

// func minimumTotal(triangle [][]int) int {
//     h := len(triangle)
//     if h == 0 {
//         return 0
//     }
//     if h == 1 {
//         return triangle[0][0]
//     }
//     dp := make([][]int, h)
//     for i := 0; i < h; i++ {
//         dp[i] = make([]int, i+1)
//     }
//     dp[0][0] = triangle[0][0]
//     min := math.MaxInt64
//     for i := 1; i < h; i++ {
//         for j := 0; j <= i; j++ {
//             if j == 0 { // left one
//                 dp[i][j] = dp[i-1][j] + triangle[i][j]
//             } else if j == i { // right one
//                 dp[i][j] = dp[i-1][j-1] +  triangle[i][j]
//             } else {
//                 tmp := dp[i-1][j-1]
//                 if dp[i-1][j] < tmp {
//                     tmp = dp[i-1][j]
//                 }
//                 dp[i][j] = tmp + triangle[i][j]
//             }
//             if i == h-1 {
//                 //fmt.Println(dp[i][j])
//                 if dp[i][j] < min {
//                     min = dp[i][j]
//                 }
//             }
//         }
//         //fmt.Println(dp)
//     }
//     return min
// }