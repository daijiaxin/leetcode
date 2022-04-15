package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp1 := 1
	dp2 := 2
	for i := 3; i <= n; i++ {
		dpi := dp1 + dp2
		dp1 = dp2
		dp2 = dpi
	}
	return dp2
}

// @lc code=end
