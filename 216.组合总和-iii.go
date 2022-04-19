package leetcode

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
var tar []int
var sum int

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	combination(k, n, 1, &res)
	return res
}

func combination(k int, n int, start int, res *[][]int) {
	if sum == n && len(tar) == k {
		tmp := []int{}
		tmp = append(tmp, tar...)
		*res = append(*res, tmp)
	}
	if sum > n || len(tar) > k {
		return
	}
	for i := start; i <= 9; i++ {
		sum += i
		tar = append(tar, i)
		combination(k, n, i+1, res)
		sum -= i
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
