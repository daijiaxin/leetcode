package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	s []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.s = append(this.s, "#")
		return "#"
	}
	this.s = append(this.s, strconv.Itoa(root.Val))
	this.serialize(root.Left)
	this.serialize(root.Right)
	return strings.Join(this.s, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	return buildTree(&s)
}

func buildTree(preorder *[]string) *TreeNode {
	if len(*preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	if (*preorder)[0] != "#" {
		root.Val, _ = strconv.Atoi((*preorder)[0])
		*preorder = (*preorder)[1:]
	} else {
		*preorder = (*preorder)[1:]
		return nil
	}
	root.Left = buildTree(preorder)
	root.Right = buildTree(preorder)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
