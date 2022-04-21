package leetcode

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	l := len(prices)
	dp0 := -prices[0]
	dp1 := 0
	for i := 1; i < l; i++ {
		d0 := max(dp0, dp1-prices[i])
		d1 := max(dp1, dp0+prices[i]-fee)
		dp0, dp1 = d0, d1
	}
	return dp1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
