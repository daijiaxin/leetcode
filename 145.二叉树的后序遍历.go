package leetcode

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	s := []*TreeNode{}
	var p *TreeNode
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		if root.Right == nil || root.Right == p {
			res = append(res, root.Val)
			p = root
			root = nil
		} else {
			s = append(s, root)
			root = root.Right
		}
	}
	return res
}

// @lc code=end
