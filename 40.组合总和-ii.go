package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
var tar []int
var sum int

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	combination(candidates, 0, target, &res)
	return res
}

func combination(candidates []int, start int, target int, res *[][]int) {
	if sum == target {
		tmp := []int{}
		tmp = append(tmp, tar...)
		*res = append(*res, tmp)
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i-1] == candidates[i] {
			continue
		}
		sum += candidates[i]
		tar = append(tar, candidates[i])
		combination(candidates, i+1, target, res)
		sum -= candidates[i]
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
