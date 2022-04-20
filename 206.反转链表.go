package leetcode

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		a := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return a
	}
	return head
}

// @lc code=end
