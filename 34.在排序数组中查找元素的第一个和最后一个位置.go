package leetcode

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left := -1
	right := -1
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			break
		}
		i := (start + end) / 2
		if nums[i] == target {
			left = i
			right = i
			ll := searchlr(nums, start, i-1, target, true)
			rr := searchlr(nums, i+1, end, target, false)
			if ll != -1 {
				left = ll
			} else {
				left = i
			}
			if rr != -1 {
				right = rr
			} else {
				right = i
			}
			break
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return []int{left, right}
}

func searchlr(nums []int, start int, end int, target int, left bool) int {
	for {
		if start > end {
			break
		}
		i := (start + end) / 2
		if nums[i] == target {
			var res int
			if left {
				res = searchlr(nums, start, end-1, target, left)
			} else {
				res = searchlr(nums, start+1, end, target, left)
			}
			if res != -1 {
				return res
			} else {
				return i
			}
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return -1
}

// @lc code=end
