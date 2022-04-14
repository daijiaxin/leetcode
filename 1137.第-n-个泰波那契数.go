package leetcode

/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp0 := 0
	dp1 := 1
	dp2 := 1
	for i := 3; i <= n; i++ {
		dpi := dp0 + dp1 + dp2
		dp0 = dp1
		dp1 = dp2
		dp2 = dpi
	}
	return dp2
}

// @lc code=end
