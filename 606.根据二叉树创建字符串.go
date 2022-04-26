package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func tree2str(root *TreeNode) string {
	var s strings.Builder
	if root == nil {
		return s.String()
	}
	s.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil && root.Right != nil {
		s.WriteString("(")
		s.WriteString(tree2str(root.Left))
		s.WriteString(")(")
		s.WriteString(tree2str(root.Right))
		s.WriteString(")")
	} else if root.Left != nil && root.Right == nil {
		s.WriteString("(")
		s.WriteString(tree2str(root.Left))
		s.WriteString(")")
	} else if root.Left == nil && root.Right != nil {
		s.WriteString("()(")
		s.WriteString(tree2str(root.Right))
		s.WriteString(")")
	}
	return s.String()
}

// @lc code=end
