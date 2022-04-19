package leetcode

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for x, boa := range board {
		for y, b := range boa {
			if b == '.' {
				continue
			}
			if isPoint(board, x, y) == false {
				return false
			}
		}
	}
	return true
}

func isPoint(board [][]byte, x int, y int) bool {
	xx := x / 3 * 3
	yy := y / 3 * 3
	for i := 0; i < 9; i++ {
		if board[i][y] == board[x][y] && i != x {
			return false
		}
		if board[x][i] == board[x][y] && i != y {
			return false
		}
		if board[xx][yy] == board[x][y] && x != xx && y != yy {
			return false
		} else {
			yy++
			if yy%3 == 0 {
				yy = y / 3 * 3
				xx++
			}
		}
	}
	return true
}

// @lc code=end
