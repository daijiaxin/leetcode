package leetcode

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root}
	var top *TreeNode
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		if top != nil && (r.Left.Val >= top.Val || r.Right.Val >= top.Val) {
			return false
		}
		if r.Left.Val >= r.Val || r.Right.Val <= r.Val {
			return false
		}
		top = r

		stack = append(stack, r.Left)
	}
	return false
}

// @lc code=end
