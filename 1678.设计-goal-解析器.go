package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=1678 lang=golang
 *
 * [1678] 设计 Goal 解析器
 */

// @lc code=start
func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(command, "()", "o")
}

// @lc code=end
