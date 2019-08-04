package dp_greedy

import "math"

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

Example 1:

Input: [2,4,1], k = 2
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
Example 2:

Input: [3,2,6,5,0,3], k = 2
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
             Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
 */

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k == 0 {
		return 0
	}
	if k > n/2 {
		k = n/2
	}
	purchaseCosts := make([]int, k)
	sellGain := make([]int, k)
	for i := 0 ; i < k; i++ {
		purchaseCosts[i] = math.MaxInt64
	}
	for i := 0; i < n; i++ {
		currentPrice := prices[i]
		var currentPurchaseCost int
		for j:= 0; j < k; j++ {
			if j == 0 {
				currentPurchaseCost = currentPrice
			} else {
				currentPurchaseCost = currentPrice - sellGain[j-1]
			}
			if currentPurchaseCost < purchaseCosts[j] {
				purchaseCosts[j] = currentPurchaseCost
			}
			currentSellGain := currentPrice - purchaseCosts[j]
			if currentSellGain > sellGain[j] {
				sellGain[j] = currentSellGain
			}
		}
	}
	return sellGain[k-1]
}