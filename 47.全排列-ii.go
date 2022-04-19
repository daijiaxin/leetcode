package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
var tar []int

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	permute(nums, &res)
	return res
}

func permute(nums []int, res *[][]int) {
	if len(nums) == 0 {
		tmp := []int{}
		tmp = append(tmp, tar...)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := []int{}
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		tar = append(tar, nums[i])
		permute(tmp, res)
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
