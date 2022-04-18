package leetcode

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, mag := range magazine {
		m[mag]++
	}
	for _, ran := range ransomNote {
		if m[ran] == 0 {
			return false
		}
		m[ran]--
	}
	return true
}

// @lc code=end
