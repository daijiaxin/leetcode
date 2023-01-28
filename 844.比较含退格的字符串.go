package leetcode

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	var ss []byte
	var tt []byte
	ls := len(s)
	lt := len(t)
	for i := 0; i < ls; i++ {
		if s[i] == '#' {
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
		} else {
			ss = append(ss, s[i])
		}
	}
	for i := 0; i < lt; i++ {
		if t[i] == '#' {
			if len(tt) > 0 {
				tt = tt[:len(tt)-1]
			}
		} else {
			tt = append(tt, t[i])
		}
	}
	return string(ss) == string(tt)
}

// @lc code=end
