package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	binaryTreePath(root, []string{}, &res)
	return res
}

func binaryTreePath(root *TreeNode, path []string, res *[]string) {
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(path, "->"))
		return
	}
	if root.Left != nil {
		binaryTreePath(root.Left, path, res)
	}
	if root.Right != nil {
		binaryTreePath(root.Right, path, res)
	}
}

// @lc code=end
