package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

// @lc code=start
func largestPerimeter(nums []int) int {
	result := 0
	count := len(nums)
	sort.Ints(nums)
	for i := count - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			diff := nums[i] - nums[j]
			for k := j - 1; k >= 0; k-- {
				if nums[j] <= diff {
					break
				}
				sum := checkLargestPerimeter(nums[i], nums[j], nums[k])
				if sum != 0 {
					return sum
				}
			}
		}
	}
	return result
}

func checkLargestPerimeter(a int, b int, c int) int {
	if a+b <= c {
		return 0
	}
	if a+c <= b {
		return 0
	}
	if b+c <= a {
		return 0
	}
	return a + b + c
}

// @lc code=end
