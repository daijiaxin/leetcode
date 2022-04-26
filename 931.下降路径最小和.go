package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				matrix[i][j] += mins(matrix[i-1][j], matrix[i-1][j+1])
				continue
			}
			if j == n-1 {
				matrix[i][j] += mins(matrix[i-1][j-1], matrix[i-1][j])
				continue
			}
			matrix[i][j] += mins(matrix[i-1][j-1], matrix[i-1][j], matrix[i-1][j+1])
		}
	}
	return mins(matrix[m-1]...)
}

func mins(nums ...int) int {
	min := math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// @lc code=end
