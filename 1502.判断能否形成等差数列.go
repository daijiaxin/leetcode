package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1502 lang=golang
 *
 * [1502] 判断能否形成等差数列
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}

// @lc code=end
