package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	preMax, preMin, ans := 1, 1, math.MinInt
	for _, num := range nums {
		preMax, preMin = max(preMax*num, preMin*num, num), min(preMax*num, preMin*num, num)
		ans = max(preMax, ans)
	}
	return ans
}

func max(vals ...int) int {
	ans := math.MinInt
	for _, v := range vals {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func min(vals ...int) int {
	ans := math.MaxInt
	for _, v := range vals {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end
