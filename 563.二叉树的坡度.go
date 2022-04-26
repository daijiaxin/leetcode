package leetcode

/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	var titl int
	sumTree(root, &titl)
	return titl
}

func sumTree(root *TreeNode, titl *int) int {
	sum := 0
	if root == nil {
		return sum
	}
	lsum := sumTree(root.Left, titl)
	rsum := sumTree(root.Right, titl)
	ti := lsum - rsum
	if ti < 0 {
		ti = -ti
	}
	*titl += ti
	return sum + lsum + rsum + root.Val
}

// @lc code=end
