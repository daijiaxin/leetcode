package leetcode

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	traverse(root, &max)
	return max
}

func traverse(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left, max)
	right := traverse(root.Right, max)
	d := left + right
	if d > *max {
		*max = d
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// @lc code=end
