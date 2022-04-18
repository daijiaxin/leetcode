package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	count := 1
	for i, num := range nums {
		if count <= i {
			return false
		}
		max := i + num + 1
		if max > count {
			count = max
		}
	}
	return true
}

// @lc code=end
