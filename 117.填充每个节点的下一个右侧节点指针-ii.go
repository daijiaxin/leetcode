package leetcode

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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
	s := []*Node{root}
	t := []*Node{}
	for len(s) > 0 || len(t) > 0 {
		if len(s) == 0 {
			s = append(s, t...)
			t = []*Node{}
		}
		n := s[0]
		s = s[1:]
		if len(s) > 0 {
			n.Next = s[0]
		}
		if n.Left != nil {
			t = append(t, n.Left)
		}
		if n.Right != nil {
			t = append(t, n.Right)
		}
	}
	return root
}

// @lc code=end
