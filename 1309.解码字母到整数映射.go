package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=1309 lang=golang
 *
 * [1309] 解码字母到整数映射
 */

// @lc code=start
func freqAlphabets(s string) string {
	n, ans := len(s), []rune{}
	for i := 0; i < n; i++ {
		if i+2 < n && s[i+2] == '#' {
			ans = append(ans, get(s[i:i+2]))
			i += 2
		} else {
			ans = append(ans, get(s[i:i+1]))
		}
	}
	return string(ans)
}

func get(s string) rune {
	r, _ := strconv.Atoi(s)
	return rune(r + 96)
}

// @lc code=end
