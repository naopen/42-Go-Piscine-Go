/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   eightqueens.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 02:39:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 02:46:20 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func isSafe(board [8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		// Check vertical
		if board[i] == col {
			return false
		}
		// Check diagonals
		if ft_abs(row-i) == ft_abs(col-board[i]) {
			return false
		}
	}
	return true
}

func solveNQUtil(board [8]int, row int) {
	if row == 8 {
		printSolution(board)
		return
	}

	for col := 0; col < 8; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveNQUtil(board, row+1)
		}
	}
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		ft.PrintRune(rune(board[i] + '1'))
	}
	ft.PrintRune('\n')
}

func ft_abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// EightQueens solves the Eight Queens puzzle
func EightQueens() {
	var board [8]int
	solveNQUtil(board, 0)
}