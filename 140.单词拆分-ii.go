package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	var res []string
	wordBreaks(s, wordDict, []string{}, &res)
	return res
}

func wordBreaks(s string, wordDict []string, path []string, res *[]string) {
	l := len(s)
	for _, word := range wordDict {
		if len(word) > l {
			continue
		}
		if s == word {
			path = append(path, word)
			*res = append(*res, strings.Join(path, " "))
			path = path[:len(path)-1]
			continue
		}
		if s[:len(word)] == word {
			path = append(path, word)
			wordBreaks(s[len(word):], wordDict, path, res)
			path = path[:len(path)-1]
		}
	}
}

// @lc code=end
