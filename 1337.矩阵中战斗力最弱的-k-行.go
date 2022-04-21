package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 矩阵中战斗力最弱的 K 行
 */

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
	n := len(mat[0])
	count := [][]int{}
	for m, mm := range mat {
		c := sort.Search(n, func(i int) bool { return mm[i] < 1 })
		count = append(count, []int{c, m})
	}
	sort.Slice(count, func(i, j int) bool {
		if count[i][0] < count[j][0] {
			return true
		} else if count[i][0] == count[j][0] {
			if count[i][1] < count[j][1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, count[i][1])
	}
	return res
}

// @lc code=end
