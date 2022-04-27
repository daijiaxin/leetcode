package leetcode

/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	smat := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := matrix[i][j]
			if i > 0 && j > 0 {
				sum -= smat[i-1][j-1]
			}
			if i > 0 {
				sum += smat[i-1][j]
			}
			if j > 0 {
				sum += smat[i][j-1]
			}
			smat[i] = append(smat[i], sum)
		}
	}
	return NumMatrix{sum: smat}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.sum[row2][col2]
	if row1 > 0 && col1 > 0 {
		sum += this.sum[row1-1][col1-1]
	}
	if row1 > 0 {
		sum -= this.sum[row1-1][col2]
	}
	if col1 > 0 {
		sum -= this.sum[row2][col1-1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
