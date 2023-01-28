package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			continue
		}
		return len(ss[i])
	}
	return 0
}

// @lc code=end
