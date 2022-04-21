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
	a []string
	b []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.a = append(this.a, strconv.Itoa(root.Val))
	this.serialize(root.Left)
	this.b = append(this.b, strconv.Itoa(root.Val))
	this.serialize(root.Right)
	s := strings.Join(this.a, ",") + "|" + strings.Join(this.b, ",")
	return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	d := strings.Split(data, "|")
	this.a = strings.Split(d[0], ",")
	this.b = strings.Split(d[1], ",")
	return buildTree(this.a, this.b)
}

func buildTree(preorder []string, inorder []string) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := 0
	for i, in := range inorder {
		if in == preorder[index] {
			index = i
			break
		}
	}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(preorder[0])
	root.Left = buildTree(preorder[1:], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
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
