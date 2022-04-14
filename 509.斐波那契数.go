package leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp0 := 0
	dp1 := 1
	for i := 2; i <= n; i++ {
		dpi := dp1 + dp0
		dp0 = dp1
		dp1 = dpi
	}
	return dp1
}

// @lc code=end
