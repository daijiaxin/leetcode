package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	start := 0
	end := len(nums) - 1
	num := len(nums) - 1
	for {
		if num < 0 {
			break
		}
		snum := int(math.Abs(float64(nums[start])))
		enum := int(math.Abs(float64(nums[end])))
		if snum > enum {
			res[num] = snum * snum
			start++
		} else {
			res[num] = enum * enum
			end--
		}
		num--
	}
	return res
}

// @lc code=end
