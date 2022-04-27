package leetcode

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	sum := 0
	base := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		sum += base * int(columnTitle[i]-'A'+1)
		base *= 26
	}
	return sum
}

// @lc code=end
