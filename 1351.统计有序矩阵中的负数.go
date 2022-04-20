package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	x := m * n
	mm := sort.Search(m, func(i int) bool { return grid[i][0] < 0 })
	for i := 0; i < mm; i++ {
		nn := sort.Search(n, func(j int) bool { return grid[i][j] < 0 })
		x -= nn
	}
	return x
}

// @lc code=end
