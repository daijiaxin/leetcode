package leetcode

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	lf := []*Node{}
	i := 1
	c := 1
	lf = append(lf, root)
	for len(lf) > 0 {
		r := lf[0]
		lf = lf[1:]
		if c == i {
			i *= 2
			c = 1
			r.Next = nil
		} else {
			c++
			r.Next = lf[0]
		}
		if r.Left != nil && r.Right != nil {
			lf = append(lf, r.Left)
			lf = append(lf, r.Right)
		}
	}
	return root
}

// @lc code=end
