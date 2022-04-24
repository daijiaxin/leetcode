package leetcode

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	s := []*TreeNode{root}
	for len(s) > 0 {
		t := s[len(s)-1]
		s = s[:len(s)-1]
		if t == nil {
			continue
		}
		if t.Left != nil && t.Left.Left == nil && t.Left.Right == nil {
			res += t.Left.Val
		}
		s = append(s, t.Left)
		s = append(s, t.Right)
	}
	return res
}

// @lc code=end
