package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		switch v {
		case "+":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b+a)
			break
		case "-":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b-a)
			break
		case "*":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b*a)
			break
		case "/":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b/a)
			break
		default:
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		}
	}
	return stack[len(stack)-1]
}

// @lc code=end
