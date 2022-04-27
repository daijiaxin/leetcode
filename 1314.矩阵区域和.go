package leetcode

/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func matrixBlockSum(mat [][]int, k int) [][]int {
	m := len(mat)
	n := len(mat[0])
	smat := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := mat[i][j]
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
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mini := i - k - 1
			minj := j - k - 1
			maxi := i + k
			maxj := j + k
			if maxi >= m {
				maxi = m - 1
			}
			if maxj >= n {
				maxj = n - 1
			}
			num := smat[maxi][maxj]
			if mini >= 0 {
				num -= smat[mini][maxj]
			}
			if minj >= 0 {
				num -= smat[maxi][minj]
			}
			if mini >= 0 && minj >= 0 {
				num += smat[mini][minj]
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}

// @lc code=end
