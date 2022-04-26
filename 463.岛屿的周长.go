package leetcode

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	ml := len(grid) - 1
	nl := len(grid[0]) - 1
	res := 0
	for m, gr := range grid {
		for n, g := range gr {
			if g == 0 {
				continue
			}
			if m == 0 || grid[m-1][n] == 0 {
				res++
			}
			if m == ml || grid[m+1][n] == 0 {
				res++
			}
			if n == 0 || grid[m][n-1] == 0 {
				res++
			}
			if n == nl || grid[m][n+1] == 0 {
				res++
			}
		}
	}
	return res
}

// @lc code=end
