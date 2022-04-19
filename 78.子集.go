package leetcode

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start

var tar []int

func subsets(nums []int) [][]int {
	var res [][]int
	subset(nums, 0, &res)
	return res
}

func subset(nums []int, start int, res *[][]int) {
	tmp := []int{}
	tmp = append(tmp, tar...)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		tar = append(tar, nums[i])
		subset(nums, i+1, res)
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
