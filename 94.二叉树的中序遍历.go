package leetcode

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	if root == nil {
		return r
	}
	r = append(r, inorderTraversal(root.Left)...)
	r = append(r, root.Val)
	r = append(r, inorderTraversal(root.Right)...)
	return r
}

// @lc code=end
