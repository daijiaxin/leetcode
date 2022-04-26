package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
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
func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt
	pr := &TreeNode{Val: -1}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pr.Val != -1 {
			d := root.Val - pr.Val
			if d < min {
				min = d
			}
		}
		pr = root
		root = root.Right
	}
	return min
}

// @lc code=end
