package leetcode

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
var tar []int

func combine(n int, k int) [][]int {
	var res [][]int
	combin(n, k, 1, &res)
	return res
}

func combin(n int, k int, start int, res *[][]int) {
	if len(tar) == k {
		tmp := []int{}
		tmp = append(tmp, tar...)
		*res = append(*res, tmp)
		return
	}
	if len(tar) > k {
		return
	}
	for i := start; i <= n; i++ {
		tar = append(tar, i)
		combin(n, k, i+1, res)
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
