package leetcode

import "math"

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
	i := math.MinInt
	return getBst(root, &i)
}

func getBst(root *TreeNode, last *int) bool {
	if root == nil {
		return true
	}
	if !getBst(root.Left, last) {
		return false
	}
	if root.Val <= *last {
		return false
	} else {
		*last = root.Val
	}
	if !getBst(root.Right, last) {
		return false
	}
	return true
}

// @lc code=end
