package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	lenght := len(s)
	for i := 0; i < lenght/2; i++ {
		a := s[i]
		s[i] = s[lenght-1-i]
		s[lenght-1-i] = a
	}
}

// @lc code=end
