package dp_greedy

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
 */

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := 0
	boughtPrice := -1
	for i:= 0; i < n-1; i++ {
		if prices[i+1] >= prices[i] {
			if boughtPrice > -1 { // 已经买了，但是还在涨，所以不管了
				continue
			} else {
				boughtPrice = prices[i] // 没买 赶紧买
			}
		} else {
			if boughtPrice > -1 { // 买了，然后后面会跌，赶紧卖
				maxProfit += prices[i] - boughtPrice
				boughtPrice = -1
			} else {
				continue // 还没买，那还行
			}
		}
		//fmt.Println(boughtPrice, maxProfit)
	}
	//fmt.Println(boughtPrice, prices[n-1])
	if boughtPrice > -1 && prices[n-1] > boughtPrice {
		maxProfit += prices[n-1] - boughtPrice
	}
	return maxProfit
}

// 最简单的办法就是都加起来
/*
class Solution {
    public int maxProfit(int[] prices) {
        int maxprofit = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] > prices[i - 1])
                maxprofit += prices[i] - prices[i - 1];
        }
        return maxprofit;
    }
}
 */