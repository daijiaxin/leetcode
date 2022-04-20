package leetcode

/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
func maxScoreSightseeingPair(values []int) int {
	max := 0
	dp := values[0]
	for i := 1; i < len(values); i++ {
		value := values[i] - i + dp
		if value > max {
			max = value
		}
		if values[i]+i > dp {
			dp = values[i] + i
		}
	}
	return max
}

// @lc code=end
