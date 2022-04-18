package leetcode

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			res = append(res, []int{1})
		} else if i == 2 {
			res = append(res, []int{1, 1})
		} else {
			rr := []int{}
			rr = append(rr, 1)
			for j := 1; j < i-1; j++ {
				rr = append(rr, res[i-2][j-1]+res[i-2][j])
			}
			rr = append(rr, 1)
			res = append(res, rr)
		}
	}
	return res
}

// @lc code=end
