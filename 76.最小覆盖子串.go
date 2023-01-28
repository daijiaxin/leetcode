package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	m := map[byte]int{}
	w := map[byte]int{}
	ls := len(s)
	lt := len(t)
	for i := 0; i < lt; i++ {
		m[t[i]]++
	}
	left := 0
	right := 0
	valid := 0
	start := 0
	l := math.MaxInt
	for right < ls {
		c := s[right]
		right++
		if m[c] > 0 {
			w[c]++
			if m[c] == w[c] {
				valid++
			}
		}
		for valid == len(m) {
			if right-left < l {
				start = left
				l = right - left
			}
			d := s[left]
			left++
			if m[d] > 0 {
				if w[d] == m[d] {
					valid--
				}
				w[d]--
			}
		}
	}
	if l == math.MaxInt {
		return ""
	}
	return s[start : start+l]
}

// @lc code=end
