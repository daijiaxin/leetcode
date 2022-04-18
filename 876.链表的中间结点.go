package leetcode

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
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
func middleNode(head *ListNode) *ListNode {
	count := 0
	ff := head
	for ff != nil {
		count++
		ff = ff.Next
	}
	cc := count / 2
	for i := 0; i < cc; i++ {
		head = head.Next
	}
	return head
}

// @lc code=end
