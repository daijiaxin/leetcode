package leetcode

/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}
	index := 0
	for i, n := range postorder {
		if n == preorder[1] {
			index = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:index+2], postorder[:index])
	if index+2 < len(preorder) {
		root.Right = constructFromPrePost(preorder[index+2:], postorder[index+1:len(postorder)-1])
	}
	return root
}

// @lc code=end
