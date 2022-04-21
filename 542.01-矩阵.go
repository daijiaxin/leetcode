package leetcode

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	ml := len(mat)
	nl := len(mat[0])
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	l := [][]int{}
	b := make([][]bool, ml)
	for m, ma := range mat {
		b[m] = make([]bool, nl)
		for n, a := range ma {
			if a == 0 {
				l = append(l, []int{m, n})
				b[m][n] = true
			}
		}
	}
	for len(l) > 0 {
		m := l[0][0]
		n := l[0][1]
		l = l[1:]
		for i := 0; i < 4; i++ {
			mm := m + dir[i][0]
			nn := n + dir[i][1]
			if mm >= 0 && mm < ml && nn >= 0 && nn < nl && !b[mm][nn] {
				mat[mm][nn] = mat[m][n] + 1
				l = append(l, []int{mm, nn})
				b[mm][nn] = true
			}
		}
	}
	return mat
}

// @lc code=end
