package leetcode

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for m, gi := range grid {
		for n := range gi {
			c := findArea(grid, m, n)
			if c > max {
				max = c
			}
		}
	}
	return max
}

func findArea(grid [][]int, m int, n int) int {
	mn := len(grid)
	nn := len(grid[0])
	if m < 0 || m >= mn || n < 0 || n >= nn || grid[m][n] == 0 {
		return 0
	}
	grid[m][n] = 0
	count := 1
	count += findArea(grid, m-1, n)
	count += findArea(grid, m+1, n)
	count += findArea(grid, m, n-1)
	count += findArea(grid, m, n+1)
	return count
}

// @lc code=end
