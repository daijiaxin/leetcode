package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	sn := strconv.Itoa(n)
	x := 1
	s := 0
	for _, c := range sn {
		ic, _ := strconv.Atoi(string(c))
		x = x * ic
		s = s + ic
	}
	return x - s
}

// @lc code=end
