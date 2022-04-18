package leetcode

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
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

func preorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}
	st := []*Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
	}
	return ans
}

// @lc code=end
