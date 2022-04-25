package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return 0
	}
	res := 0
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			e := search(nums, nums[i]+nums[j])
			if e > j {
				res += e - j
			}
		}
	}
	return res
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		i := (start + end) / 2
		if nums[i] >= target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return end
}

// @lc code=end
