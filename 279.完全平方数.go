package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	num := int(math.Sqrt(float64(n)))
	if num*num == n {
		return 1
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n
	}
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
