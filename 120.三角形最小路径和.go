package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	dp := []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		t := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				t[j] = dp[j] + triangle[i][j]
			} else if j == i {
				t[j] = dp[j-1] + triangle[i][j]
			} else {
				t[j] = dp[j] + triangle[i][j]
				if dp[j-1]+triangle[i][j] < t[j] {
					t[j] = dp[j-1] + triangle[i][j]
				}
			}
		}
		dp = t
	}
	res := math.MaxInt
	for _, d := range dp {
		if d < res {
			res = d
		}
	}
	return res
}

// @lc code=end
