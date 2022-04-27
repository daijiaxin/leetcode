package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	nums := []int{}
	for _, env := range envelopes {
		nums = append(nums, env[1])
	}
	return lengthOfLIS(nums)
}

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
