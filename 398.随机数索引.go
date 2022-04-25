package leetcode

/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
type Solution struct {
	m map[int]SolutionNum
}

type SolutionNum struct {
	index []int
	last  int
}

func Constructor(nums []int) Solution {
	m := map[int]SolutionNum{}
	for i, n := range nums {
		mn := m[n]
		mn.index = append(mn.index, i)
		m[n] = mn
	}
	return Solution{m}
}

func (this *Solution) Pick(target int) int {
	mn := this.m[target]
	res := mn.index[mn.last]
	mn.last++
	if mn.last == len(mn.index) {
		mn.last = 0
	}
	this.m[target] = mn
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
