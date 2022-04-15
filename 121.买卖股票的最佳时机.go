package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	low := 100000
	max := 0
	for i := 0; i < n; i++ {
		if prices[i] < low {
			low = prices[i]
		}
		m := prices[i] - low
		if m > max {
			max = m
		}
	}
	return max
}

// @lc code=end
