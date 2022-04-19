package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
var tar []string
var count int

func getPermutation(n int, k int) string {
	var res [][]string
	var nums []string
	cc := 1
	for i := 1; i <= n; i++ {
		nums = append(nums, fmt.Sprintf("%d", i))
		cc *= i
	}
	start := (k - 1) / (cc / n)
	if n > 3 {
		getPremut(nums, start, &res)
		return strings.Join(res[(k-1)%(cc/n)], "")
	} else {
		getPremut(nums, 0, &res)
		return strings.Join(res[k-1], "")
	}

}

func getPremut(nums []string, start int, res *[][]string) {
	if len(nums) == 0 {
		tmp := []string{}
		tmp = append(tmp, tar...)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := []string{}
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		tar = append(tar, nums[i])
		getPremut(tmp, 0, res)
		tar = tar[:len(tar)-1]
	}
}

// @lc code=end
