package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt
	}
	if len(nums) == 1 {
		return nums[0]
	}
	start := 0
	end := len(nums) - 1
	if nums[start] < nums[end] {
		return nums[start]
	}
	i := (start + end) / 2
	l := findMin(nums[:i+1])
	r := findMin(nums[i+1:])
	if l < r {
		return l
	}
	return r
}

// @lc code=end
