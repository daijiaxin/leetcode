package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	mm := sort.Search(m, func(i int) bool { return matrix[i][0] >= target })
	if mm != m && matrix[mm][0] == target {
		return true
	}
	if mm == 0 {
		return false
	}
	nn := sort.Search(n, func(i int) bool { return matrix[mm-1][i] >= target })
	if nn == n {
		return false
	} else {
		if matrix[mm-1][nn] == target {
			return true
		} else if nn > 0 && matrix[mm-1][nn-1] == target {
			return true
		}
	}
	return false
}

// @lc code=end
