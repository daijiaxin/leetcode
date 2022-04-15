package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	result := 0
	for _, a1 := range arr1 {
		add := true
		for _, a2 := range arr2 {
			if math.Abs(float64(a1-a2)) <= float64(d) {
				add = false
				break
			}
		}
		if add {
			result++
		}
	}
	return result
}

// @lc code=end
