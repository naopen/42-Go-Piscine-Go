/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:54:24 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/10 18:00:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintNumber prints a single digit number.
func PrintNumber(n int) {
	ft.PrintRune(rune(n) + '0')
}

// PrintComb2 prints all possible combinations of two different two-digit numbers in ascending order.
func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			PrintNumber(i / 10)
			PrintNumber(i % 10)
			ft.PrintRune(' ')
			PrintNumber(j / 10)
			PrintNumber(j % 10)
			if i != 98 || j != 99 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
