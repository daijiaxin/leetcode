package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		are := min(height[l], height[r]) * (r - l)
		if are > res {
			res = are
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
