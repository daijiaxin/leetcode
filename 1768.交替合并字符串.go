package leetcode

/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] 交替合并字符串
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	w1 := 0
	w2 := 0
	l1 := len(word1)
	l2 := len(word2)
	w := true
	for w1 < l1 || w2 < l2 {
		if w && w1 < l1 {
			res = append(res, word1[w1])
			w1++
			if w2 < l2 {
				w = false
			}
		}
		if !w && w2 < l2 {
			res = append(res, word2[w2])
			w2++
			if w1 < l1 {
				w = true
			}
		}
	}
	return string(res)
}

// @lc code=end
