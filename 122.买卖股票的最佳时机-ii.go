package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	sum := 0
	l := len(prices)
	start := prices[0]
	for i := 1; i < l; i++ {
		if prices[i] > start {
			sum += prices[i] - start
		}
		start = prices[i]
	}
	return sum
}

// @lc code=end
