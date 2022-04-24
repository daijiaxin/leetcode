package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	res := 0
	length := len(height)
	l := make([]int, length)
	r := make([]int, length)
	l[0] = height[0]
	r[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		l[i] = max(l[i-1], height[i])
	}
	for i := length - 2; i >= 0; i-- {
		r[i] = max(r[i+1], height[i])
	}
	for i := 1; i < length-1; i++ {
		res += min(l[i], r[i]) - height[i]
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
