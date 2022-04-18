package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	res := []byte{}
	for jj, sss := range ss {
		for i := len(sss) - 1; i >= 0; i-- {
			res = append(res, sss[i])
		}
		if jj != len(ss)-1 {
			res = append(res, ' ')
		}
	}
	return string(res)
}

// @lc code=end
