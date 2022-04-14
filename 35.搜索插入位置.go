package leetcode

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			return start
		}
		i := (start + end) / 2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			start = i + 1
		} else {
			end = i - 1
		}
	}
}

// @lc code=end
