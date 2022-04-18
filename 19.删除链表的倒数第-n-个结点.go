package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hh := head
	count := 0
	for hh != nil {
		count++
		hh = hh.Next
	}
	ff := count - n
	hh = head
	for i := 0; i < ff-1; i++ {
		hh = hh.Next
	}
	if ff == 0 {
		return head.Next
	}
	if head.Next == nil {
		return head
	}
	hh.Next = hh.Next.Next
	return head
}

// @lc code=end
