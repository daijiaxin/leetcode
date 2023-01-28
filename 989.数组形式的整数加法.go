package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(num []int, k int) []int {
	kl := strings.Split(strconv.Itoa(k), "")
	ln := len(num)
	lk := len(kl)
	add := 0
	index := 1
	for ln-index >= 0 || lk-index >= 0 {
		in := ln - index
		ik := lk - index
		var kn int
		if ik >= 0 {
			kn, _ = strconv.Atoi(kl[ik])
		}
		if in < 0 {
			ss := kn + add
			num = append([]int{ss % 10}, num...)
			add = ss / 10
			ln++
		} else {
			ss := num[in] + kn + add
			num[in] = ss % 10
			add = ss / 10
		}
		index++
	}
	if add == 1 {
		num = append([]int{1}, num...)
	}
	return num
}

// @lc code=end
