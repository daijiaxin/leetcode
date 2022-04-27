package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	return calcudp(dungeon, &dp, 0, 0)
}

func calcudp(dungeon [][]int, dp *[][]int, i int, j int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	if i == m-1 && j == n-1 {
		if dungeon[i][j] < 0 {
			return -dungeon[i][j] + 1
		} else {
			return 1
		}
	}
	if i == m || j == n {
		return math.MaxInt
	}
	if (*dp)[i][j] > 0 {
		return (*dp)[i][j]
	}
	res := min(calcudp(dungeon, dp, i, j+1), calcudp(dungeon, dp, i+1, j)) - dungeon[i][j]
	if res > 0 {
		(*dp)[i][j] = res
	} else {
		(*dp)[i][j] = 1
	}
	return (*dp)[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
