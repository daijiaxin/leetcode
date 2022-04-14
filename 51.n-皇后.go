package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	tmp := make([][]string, n)
	for i := range tmp {
		for j := 0; j < n; j++ {
			tmp[i] = append(tmp[i], ".")
		}
	}
	solve(tmp, 0, &res)
	return res
}

func solve(tmp [][]string, r int, res *[][]string) {
	if len(tmp) == r {
		var re []string
		for _, tm := range tmp {
			re = append(re, strings.Join(tm, ""))
		}
		*res = append(*res, re)
		return
	}
	for c := 0; c < len(tmp); c++ {
		if !isOk(tmp, r, c) {
			continue
		}
		tmp[r][c] = "Q"
		solve(tmp, r+1, res)
		tmp[r][c] = "."
	}
}

func isOk(tmp [][]string, r int, c int) bool {
	for i := 0; i < len(tmp); i++ {
		if tmp[i][c] == "Q" {
			return false
		}
	}
	il := r - 1
	jl := c - 1
	for {
		if il < 0 || jl < 0 {
			break
		}
		if tmp[il][jl] == "Q" {
			return false
		}
		il--
		jl--
	}
	ir := r - 1
	jr := c + 1
	for {
		if ir < 0 || jr >= len(tmp) {
			break
		}
		if tmp[ir][jr] == "Q" {
			return false
		}
		ir--
		jr++
	}
	return true
}

// @lc code=end
