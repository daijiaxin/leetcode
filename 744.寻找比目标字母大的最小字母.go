package leetcode

/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] < target {
		return letters[0]
	}
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}

// @lc code=end
