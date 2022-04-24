package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
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
func getDecimalValue(head *ListNode) int {
	b := []string{}
	for head != nil {
		b = append(b, strconv.Itoa(head.Val))
		head = head.Next
	}
	i, _ := strconv.ParseInt(strings.Join(b, ""), 2, 64)
	return int(i)
}

// @lc code=end
