package leetcode

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}
	res = append(res, root.Val)
	return res
}

// @lc code=end
