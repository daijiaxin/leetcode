package leetcode

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

func guess(n int) int

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	start := 0
	end := n
	for {
		i := (start + end) / 2
		res := guess(i)
		if res == 0 {
			return i
		} else if res == 1 {
			start = i + 1
		} else {
			end = i - 1
		}
	}
}

// @lc code=end
