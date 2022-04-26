package leetcode

/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	stack := []*TreeNode{root}
	nstack := []*TreeNode{}
	var sum float64
	var count float64
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t != nil {
			sum += float64(t.Val)
			count++
			nstack = append(nstack, t.Left, t.Right)
		}
		if len(stack) == 0 {
			if count != 0 {
				res = append(res, sum/count)
				sum = 0
				count = 0
			}
			stack = append(stack, nstack...)
			nstack = []*TreeNode{}
		}
	}
	return res
}

// @lc code=end
