package leetcode

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp1 := -prices[0]
	dp2 := 0
	dp3 := 0
	for i := 1; i < l; i++ {
		d1 := max(dp1, dp3-prices[i])
		d2 := dp1 + prices[i]
		d3 := max(dp2, dp3)
		dp1, dp2, dp3 = d1, d2, d3
	}
	return max(dp2, dp3)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
