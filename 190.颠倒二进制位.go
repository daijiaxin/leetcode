package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	s := strconv.FormatUint(uint64(num), 2)
	var r strings.Builder
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		r.WriteByte(s[i])
	}
	for i := l; i < 32; i++ {
		r.WriteByte('0')
	}
	us, _ := strconv.ParseUint(r.String(), 2, 64)
	return uint32(us)
}

// @lc code=end
