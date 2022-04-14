package leetcode

/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 */

// @lc code=start
func countOdds(low int, high int) int {
	res := (high - low + 1) / 2
	if low&1 == 1 && high&1 == 1 {
		return res + 1
	}
	return res
}

// @lc code=end
