package leetcode

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	l := getList(root)
	start := 0
	end := len(l) - 1
	for start < end {
		sum := l[start] + l[end]
		if sum == k {
			return true
		} else if sum > k {
			end--
		} else {
			start++
		}
	}
	return false
}

func getList(root *TreeNode) []int {
	l := []int{}
	if root == nil {
		return l
	}
	l = append(l, getList(root.Left)...)
	l = append(l, root.Val)
	l = append(l, getList(root.Right)...)
	return l
}

// @lc code=end
