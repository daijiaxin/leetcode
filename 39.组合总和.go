package leetcode

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

var tar []int
var sum int

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
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
		sum += candidates[i]
		tar = append(tar, candidates[i])
		combination(candidates, i, target, res)
		sum -= candidates[i]
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
