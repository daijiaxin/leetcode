package leetcode

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
		t.Left = nil
		if len(stack) > 0 {
			t.Right = stack[len(stack)-1]
		}
	}
}

// @lc code=end
