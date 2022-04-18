package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	start := 0
	max := 0
	mp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if p, ok := mp[s[i]]; ok {
			if p > start {
				start = p
			}
		}
		mp[s[i]] = i + 1
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

// @lc code=end
