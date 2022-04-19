package leetcode

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	sudoku(board, 0, 0)
}

func sudoku(board [][]byte, x int, y int) bool {
	if y == 9 {
		return sudoku(board, x+1, 0)
	}
	if x == 9 {
		return true
	}
	if board[x][y] != '.' {
		return sudoku(board, x, y+1)
	}
	var i byte = '1'
	for ; i <= '9'; i++ {
		if !isPoint(board, x, y, i) {
			continue
		}
		board[x][y] = i
		if sudoku(board, x, y+1) {
			return true
		}
		board[x][y] = '.'
	}
	return false
}

func isPoint(board [][]byte, x int, y int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == n {
			return false
		}
		if board[x][i] == n {
			return false
		}
		if board[(x/3)*3+i/3][(y/3)*3+i%3] == n {
			return false
		}
	}
	return true
}

// @lc code=end
