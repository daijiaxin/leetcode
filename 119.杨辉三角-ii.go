package leetcode

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := []int{}
	for i := 1; i <= rowIndex+1; i++ {
		if i == 1 {
			res = []int{1}
		} else if i == 2 {
			res = []int{1, 1}
		} else {
			rr := []int{}
			rr = append(rr, 1)
			for j := 1; j < i-1; j++ {
				rr = append(rr, res[j-1]+res[j])
			}
			rr = append(rr, 1)
			res = rr
		}
	}
	return res
}

// @lc code=end
