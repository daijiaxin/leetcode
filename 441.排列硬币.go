package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}

// @lc code=end
