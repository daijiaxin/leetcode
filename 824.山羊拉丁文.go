package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=start
func toGoatLatin(sentence string) string {
	m := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	sen := strings.Split(sentence, " ")
	res := []string{}
	for i, s := range sen {
		t := []byte(s)
		if m[s[0]] != 1 {
			t = append(t[1:], t[0])
		}
		t = append(t, 'm', 'a')
		for j := 0; j < i+1; j++ {
			t = append(t, 'a')
		}
		res = append(res, string(t))
	}
	return strings.Join(res, " ")
}

// @lc code=end
