package leetcode

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	r := []int{}
	s := []*TreeNode{root}
	for len(s) > 0 {
		t := s[len(s)-1]
		s = s[:len(s)-1]
		if t != nil {
			r = append(r, t.Val)
			s = append(s, t.Right)
			s = append(s, t.Left)
		}
	}
	return r
}

// @lc code=end
