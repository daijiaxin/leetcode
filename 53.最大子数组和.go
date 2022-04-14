package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			dp[i] = num
		} else {
			a := dp[i-1] + num
			n := num
			if a >= n {
				dp[i] = a
			} else {
				dp[i] = n
			}
		}
	}
	res := math.MinInt
	for _, n := range dp {
		if n > res {
			res = n
		}
	}
	return res
}

// @lc code=end
