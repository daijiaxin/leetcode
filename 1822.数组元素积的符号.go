package leetcode

/*
 * @lc app=leetcode.cn id=1822 lang=golang
 *
 * [1822] 数组元素积的符号
 */

// @lc code=start
func arraySign(nums []int) int {
	x := 1
	for _, num := range nums {
		if num == 0 {
			return num
		} else if num < 0 {
			x = x * -1
		}
	}
	if x > 0 {
		return 1
	} else {
		return -1
	}
}

// @lc code=end
