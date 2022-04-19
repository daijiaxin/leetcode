package leetcode

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	t := head
	m := map[*ListNode]struct{}{}
	for t != nil {
		if _, ok := m[t]; ok {
			return t
		}
		m[t] = struct{}{}
		t = t.Next
	}
	return nil
}

// @lc code=end
