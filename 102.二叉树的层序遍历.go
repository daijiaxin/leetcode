package leetcode

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	s := []*TreeNode{root}
	n := []*TreeNode{}
	rr := []int{}
	for len(s) > 0 {
		t := s[0]
		s = s[1:]
		if t != nil {
			n = append(n, t.Left, t.Right)
			rr = append(rr, t.Val)
		}
		if len(s) == 0 {
			if len(rr) > 0 {
				res = append(res, rr)
				rr = []int{}
			}
			if len(n) > 0 {
				s = append(s, n...)
				n = []*TreeNode{}
			}
		}
	}
	return res
}

// @lc code=end
