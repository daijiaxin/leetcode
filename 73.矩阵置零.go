package leetcode

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	r := len(matrix)
	c := len(matrix[0])
	pot := [][]int{}
	for m, mat := range matrix {
		for n, ma := range mat {
			if ma == 0 {
				pot = append(pot, []int{m, n})
			}
		}
	}
	for _, pp := range pot {
		for i := 0; i < c; i++ {
			matrix[pp[0]][i] = 0
		}
		for i := 0; i < r; i++ {
			matrix[i][pp[1]] = 0
		}
	}
}

// @lc code=end
