package leetcode

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	return generateTree(1, n)
}

func generateTree(l int, r int) []*TreeNode {
	res := []*TreeNode{}
	if l > r {
		res = append(res, nil)
		return res
	}
	for i := l; i <= r; i++ {
		ll := generateTree(l, i-1)
		rl := generateTree(i+1, r)
		for _, lt := range ll {
			for _, rt := range rl {
				t := &TreeNode{
					Val:   i,
					Left:  lt,
					Right: rt,
				}
				res = append(res, t)
			}
		}
	}
	return res
}

// @lc code=end
