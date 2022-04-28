package leetcode

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
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
