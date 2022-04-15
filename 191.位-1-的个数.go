package leetcode

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	for {
		if num == 0 {
			break
		}
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

// @lc code=end
