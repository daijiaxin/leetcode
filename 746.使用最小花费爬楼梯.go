package leetcode

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	dp1 := 0
	dp2 := 0
	n := len(cost)
	for i := 2; i <= n; i++ {
		var dpi int
		tp1 := dp1 + cost[i-2]
		tp2 := dp2 + cost[i-1]
		if tp1 > tp2 {
			dpi = tp2
		} else {
			dpi = tp1
		}
		dp1 = dp2
		dp2 = dpi
	}
	return dp2
}

// @lc code=end
