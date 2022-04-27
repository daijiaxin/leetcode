package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := []int{nums[0]}
	for i := 1; i < l; i++ {
		d := sort.Search(len(dp), func(j int) bool { return dp[j] >= nums[i] })
		if d == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[d] = nums[i]
		}
	}
	return len(dp)
}

// @lc code=end
