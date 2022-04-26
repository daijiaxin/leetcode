package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1894 lang=golang
 *
 * [1894] 找到需要补充粉笔的学生编号
 */

// @lc code=start
func chalkReplacer(chalk []int, k int) int {
	l := len(chalk)
	sr := []int{}
	s := 0
	for _, v := range chalk {
		s += v
		sr = append(sr, s)
	}
	k = k % sr[len(sr)-1]
	return sort.Search(l, func(i int) bool { return sr[i] > k })
}

// @lc code=end
