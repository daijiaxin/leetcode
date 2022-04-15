package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int
	return gpermute(nums, []int{}, res)
}

func gpermute(nums []int, tmp []int, res [][]int) [][]int {
	for i, num := range nums {
		if num == math.MaxInt {
			continue
		}
		tmp = append(tmp, num)
		if len(nums) == len(tmp) {
			r := make([]int, len(tmp))
			copy(r, tmp)
			res = append(res, r)
			return res
		}
		nums[i] = math.MaxInt
		res = gpermute(nums, tmp, res)
		tmp = tmp[:len(tmp)-1]
		nums[i] = num
	}
	return res
}

// @lc code=end
