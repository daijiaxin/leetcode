package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1346 lang=golang
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i, a := range arr {
		ii := search(arr, a*2)
		if ii != -1 && ii != i {
			return true
		}
	}
	return false
}

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
