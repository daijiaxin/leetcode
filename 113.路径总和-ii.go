package leetcode

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	getPath(root, targetSum, []int{}, &res)
	return res
}

func getPath(root *TreeNode, target int, path []int, res *[][]int) {
	if root == nil {
		return

	}
	path = append(path, root.Val)
	if target == root.Val && root.Left == nil && root.Right == nil {
		t := []int{}
		t = append(t, path...)
		*res = append(*res, t)
	}
	getPath(root.Left, target-root.Val, path, res)
	getPath(root.Right, target-root.Val, path, res)
}

// @lc code=end
