package leetcode

/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(nums []int) bool {
	l := len(nums)
	if l <= 2 {
		return true
	}
	d := 0
	for i := 0; i < l-1; i++ {
		dd := nums[i] - nums[i+1]
		if dd == 0 {
			continue
		}
		if d == 0 {
			d = dd
			continue
		} else if d > 0 {
			if dd < 0 {
				return false
			}
		} else {
			if dd > 0 {
				return false
			}
		}
	}
	return true
}

// @lc code=end
