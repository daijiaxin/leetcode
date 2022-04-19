package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
var tar []int

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	subsetWd(nums, 0, &res)
	return res
}

func subsetWd(nums []int, start int, res *[][]int) {
	tmp := []int{}
	tmp = append(tmp, tar...)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tar = append(tar, nums[i])
		subsetWd(nums, i+1, res)
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
