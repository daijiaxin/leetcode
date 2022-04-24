package leetcode

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return isSymm(root.Left, root.Right)
}

func isSymm(left *TreeNode, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		return isSymm(left.Left, right.Right) && isSymm(left.Right, right.Left)
	} else if left == nil && right == nil {
		return true
	} else {
		return false
	}
}

// @lc code=end
