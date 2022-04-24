package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ci := onesCount(arr[i])
		cj := onesCount(arr[j])
		return ci < cj || ci == cj && arr[i] < arr[j]
	})
	return arr
}

func onesCount(x int) int {
	c := 0
	for ; x > 0; x /= 2 {
		c += x % 2
	}
	return c
}

// @lc code=end
