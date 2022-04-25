package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	sum := []int{0}
	s := 0
	for _, num := range nums {
		s += num
		sum = append(sum, s)
	}
	min := math.MaxInt
	for i, m := range sum {
		e := search(sum, m+target)
		if e != l+1 && e-i < min {
			min = e - i
		}
	}
	if min == math.MaxInt {
		min = 0
	}
	return min
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		i := (start + end) / 2
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return start
}

// @lc code=end
