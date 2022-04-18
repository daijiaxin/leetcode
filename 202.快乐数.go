package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	mp := make(map[int]int)
	for {
		mp[n] = 1
		sn := strconv.Itoa(n)
		sum := 0
		for _, s := range sn {
			ns, _ := strconv.Atoi(string(s))
			sum = sum + ns*ns
		}
		if sum == 1 {
			return true
		} else {
			n = sum
			if mp[n] == 1 {
				return false
			}
		}
	}
}

// @lc code=end
