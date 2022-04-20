package leetcode

/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	m := len(mat)
	if m == 1 {
		return mat[0][0]
	}
	sum := 0
	for i := 0; i < m; i++ {
		sum += mat[i][i]
		sum += mat[m-i-1][i]
	}
	if m%2 == 1 {
		sum -= mat[m/2][m/2]
	}
	return sum
}

// @lc code=end
