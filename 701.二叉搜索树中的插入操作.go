package leetcode

/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	res := root
	if root == nil {
		return &TreeNode{Val: val}
	}
	for {
		if root.Val > val && root.Left != nil {
			root = root.Left
			continue
		}
		if root.Val < val && root.Right != nil {
			root = root.Right
			continue
		}
		if root.Val > val {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Right = &TreeNode{Val: val}
		}
		break
	}
	return res
}

// @lc code=end
