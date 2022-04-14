package leetcode

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			break
		}
		i := (start + end) / 2
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return -1
}

// @lc code=end
