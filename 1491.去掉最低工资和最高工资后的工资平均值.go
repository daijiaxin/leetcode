package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=1491 lang=golang
 *
 * [1491] 去掉最低工资和最高工资后的工资平均值
 */

// @lc code=start
func average(salary []int) float64 {
	sum := 0
	max := math.MinInt
	min := math.MaxInt
	for _, s := range salary {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
		sum = sum + s
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}

// @lc code=end
