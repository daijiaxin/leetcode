package leetcode

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
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

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return maxD(root, 1)
}

func maxD(root *Node, d int) int {
	max := d
	for _, v := range root.Children {
		md := maxD(v, d+1)
		if md > max {
			max = md
		}
	}
	return max
}

// @lc code=end
