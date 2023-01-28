package leetcode

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	res := []int{}
	ls := len(s)
	lp := len(p)
	need := map[byte]int{}
	win := map[byte]int{}
	for i := 0; i < lp; i++ {
		need[p[i]]++
	}
	left := 0
	right := 0
	val := 0
	for right < ls {
		c := s[right]
		right++
		if need[c] > 0 {
			win[c]++
			if win[c] == need[c] {
				val++
			}
		}
		for right-left >= lp {
			if val == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if win[d] == need[d] {
					val--
				}
				win[d]--
			}
		}
	}
	return res
}

// @lc code=end
