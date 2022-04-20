package leetcode

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	var depth int
	traverse(root, 0, &depth)
	return depth
}

func traverse(root *TreeNode, depth int, max *int) {
	if root == nil {
		if depth > *max {
			*max = depth
		}
		return
	}
	depth++
	traverse(root.Left, depth, max)
	traverse(root.Right, depth, max)
	depth--
}

// @lc code=end
