package leetcode

/*
 * @lc app=leetcode.cn id=1646 lang=golang
 *
 * [1646] 获取生成数组中的最大值
 */

// @lc code=start
func getMaximumGenerated(n int) int {
	max := 0
	if n == 1 {
		return 1
	}
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		a := i / 2
		b := i % 2
		m := dp[a]
		if b == 1 {
			m += dp[a+1]
		}
		if m >= max {
			max = m
		}
		dp = append(dp, m)
	}
	return max
}

// @lc code=end
