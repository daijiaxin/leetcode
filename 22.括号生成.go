package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var res []string
	return gen(n, n, []string{}, res)
}

func gen(left int, right int, tmp []string, res []string) []string {
	if left == 0 && right == 0 {
		res = append(res, strings.Join(tmp, ""))
	} else if left == right {
		tmp = append(tmp, "(")
		left--
		res = gen(left, right, tmp, res)
		tmp = tmp[:len(tmp)-1]
	} else if left == 0 {
		tmp = append(tmp, ")")
		right--
		res = gen(left, right, tmp, res)
		tmp = tmp[:len(tmp)-1]
	} else {
		tmp = append(tmp, "(")
		left--
		res = gen(left, right, tmp, res)
		tmp = tmp[:len(tmp)-1]
		tmp = append(tmp, ")")
		left++
		right--
		res = gen(left, right, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
	return res
}

// @lc code=end
