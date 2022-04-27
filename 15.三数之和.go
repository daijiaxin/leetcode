package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}
	ll := len(nums)
	if ll < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < ll-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l := i + 1
		r := ll - 1
		for l < r {
			n2 := nums[l]
			n3 := nums[r]
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// @lc code=end
