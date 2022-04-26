package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return strings.Join(getTree(root1), ",") == strings.Join(getTree(root2), ",")
}

func getTree(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	res = append(res, getTree(root.Left)...)
	res = append(res, getTree(root.Right)...)
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
	}
	return res
}

// @lc code=end
