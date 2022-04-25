package leetcode

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	l := len(digits) - 1
	for i := l; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
