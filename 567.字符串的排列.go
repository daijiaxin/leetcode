package leetcode

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	var arr1, arr2 [26]int
	for i := range s1 {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}

	if arr1 == arr2 {
		return true
	}

	for i := len1; i < len2; i++ {
		arr2[s2[i]-'a']++
		arr2[s2[i-len1]-'a']--
		if arr1 == arr2 {
			return true
		}
	}

	return false
}

// @lc code=end
