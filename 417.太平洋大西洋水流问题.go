package leetcode

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	po := make([][]int, m)
	ao := make([][]int, m)
	for i := 0; i < m; i++ {
		po[i] = make([]int, n)
		ao[i] = make([]int, n)
	}
	op := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	poq := [][]int{}
	aoq := [][]int{}
	for i := 0; i < m; i++ {
		poq = append(poq, []int{i, 0})
		po[i][0] = 1
		aoq = append(aoq, []int{i, n - 1})
		ao[i][n-1] = 1
	}
	for i := 0; i < n; i++ {
		poq = append(poq, []int{0, i})
		po[0][i] = 1
		aoq = append(aoq, []int{m - 1, i})
		ao[m-1][i] = 1
	}
	for len(poq) > 0 {
		temp := poq[0]
		poq = poq[1:]
		i := temp[0]
		j := temp[1]
		h := heights[i][j]
		for k := 0; k < 4; k++ {
			ii := i + op[k][0]
			jj := j + op[k][1]
			if ii < 0 || jj < 0 || ii == m || jj == n || po[ii][jj] == 1 {
				continue
			}
			if heights[ii][jj] >= h {
				poq = append(poq, []int{ii, jj})
				po[ii][jj] = 1
			}
		}
	}
	for len(aoq) > 0 {
		temp := aoq[0]
		aoq = aoq[1:]
		i := temp[0]
		j := temp[1]
		h := heights[i][j]
		for k := 0; k < 4; k++ {
			ii := i + op[k][0]
			jj := j + op[k][1]
			if ii < 0 || jj < 0 || ii == m || jj == n || ao[ii][jj] == 1 {
				continue
			}
			if heights[ii][jj] >= h {
				aoq = append(aoq, []int{ii, jj})
				ao[ii][jj] = 1
			}
		}
	}
	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if po[i][j] == 1 && ao[i][j] == 1 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

// @lc code=end
