package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	n := []int{}
	last := math.MaxInt
	for _, num := range nums {
		if num == last || (len(n) > 1 && num == n[0]) {
			continue
		}
		n = append(n, num)
		last = num
	}
	return ssearch(n, target)
}

func ssearch(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			return false
		}
		i := (start + end) / 2
		if nums[i] == target {
			return true
		} else if nums[start] > nums[end] {
			l := ssearch(nums[:i], target)
			if l {
				return true
			}
			r := ssearch(nums[i+1:], target)
			if r {
				return true
			}
			return false
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
