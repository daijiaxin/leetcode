package leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	dp1 := 0
	dp2 := 0
	for _, num := range nums {
		dp1, dp2 = dp2, max(dp2, dp1+num)
	}
	return dp2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
