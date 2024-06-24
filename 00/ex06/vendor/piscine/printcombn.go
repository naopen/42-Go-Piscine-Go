/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcombn.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 18:16:54 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 00:33:22 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintNumber prints a single digit number.
func PrintNumber(n int) {
	ft.PrintRune(rune(n) + '0')
}

func printCombination(digits []int) {
	for i := 0; i < SliceLen(digits); i++ {
		PrintNumber(digits[i])
	}
	if digits[0] != 9-(SliceLen(digits)-1) {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func generateCombinations(digits []int, index int, start int, n int) {
	if index == n {
		printCombination(digits)
		return
	}
	for i := start; i <= 9-(n-index-1); i++ {
		digits[index] = i
		generateCombinations(digits, index+1, i+1, n)
	}
}

// PrintCombN prints all possible combinations of n different digits in ascending order.
func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var digits [10]int
	generateCombinations(digits[:n], 0, 0, n)
	ft.PrintRune('\n')
}
