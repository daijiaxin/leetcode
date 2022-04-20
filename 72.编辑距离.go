package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start

func minDistance(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = mins(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m-1][n-1]
}

func mins(i ...int) int {
	min := math.MaxInt
	for _, m := range i {
		if m < min {
			min = m
		}
	}
	return min
}

// @lc code=end
