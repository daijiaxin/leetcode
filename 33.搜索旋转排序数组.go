package leetcode

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			return -1
		}
		i := (start + end) / 2
		if nums[i] == target {
			return i
		} else if nums[start] > nums[end] {
			l := search(nums[:i], target)
			if l != -1 {
				return l
			}
			r := search(nums[i+1:], target)
			if r != -1 {
				return r + i + 1
			}
			return -1
		} else {
			if nums[i] > target {
				end = i - 1
			} else {
				start = i + 1
			}
		}
	}
}

// @lc code=end
