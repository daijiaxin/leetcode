package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
var tar []string
var count int

func getPermutation(n int, k int) string {
	var res []string
	var nums []string
	var xx []int
	cc := 1
	for i := 1; i <= n; i++ {
		nums = append(nums, fmt.Sprintf("%d", i))
		xx = append(xx, cc)
		cc *= i
	}
	k--
	for len(nums) > 0 {
		index := len(nums) - 1
		num := k / xx[index]
		k = k % xx[index]
		res = append(res, nums[num])
		nums = append(nums[:num], nums[num+1:]...)
	}
	return strings.Join(res, "")
}

// @lc code=end
