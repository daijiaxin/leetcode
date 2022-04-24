package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	l := -1
	r := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		} else if l == -1 {
			l = i
		} else {
			if i-l > r {
				r = i - l
			}
			l = i
		}
	}
	return r
}

// @lc code=end
