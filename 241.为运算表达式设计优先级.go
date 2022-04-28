package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	res := []int{}
	l := len(expression)
	for i := 0; i < l; i++ {
		c := expression[i]
		if c != '+' && c != '-' && c != '*' {
			continue
		}
		l := diffWaysToCompute(expression[0:i])
		r := diffWaysToCompute(expression[i+1:])
		for _, a := range l {
			for _, b := range r {
				if c == '+' {
					res = append(res, a+b)
				} else if c == '-' {
					res = append(res, a-b)
				} else {
					res = append(res, a*b)
				}
			}
		}
	}
	if len(res) == 0 {
		d, _ := strconv.Atoi(expression)
		res = append(res, d)
	}
	return res
}

// @lc code=end
