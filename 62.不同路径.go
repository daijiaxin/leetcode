package leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dpl := 0
			dpt := 0
			if i > 0 {
				dpt = dp[i-1][j]
			}
			if j > 0 {
				dpl = dp[i][j-1]
			}
			dp[i][j] = dpt + dpl
			if i == 0 && j == 0 {
				dp[i][j]++
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
