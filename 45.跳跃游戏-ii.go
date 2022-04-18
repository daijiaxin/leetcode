package leetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	steps := 0
	max := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if i == end {
			end = max
			steps++
		}
	}
	return steps
}

// @lc code=end
