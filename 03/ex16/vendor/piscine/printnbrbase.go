/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printnbrbase.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:35:14 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:56:54 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func isValidBase(base string) bool {
	if StrLen(base) < 2 {
		return false
	}
	for i := 0; i < StrLen(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < StrLen(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

// PrintNbrBase prints an integer in a given base.
func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		ft.PrintRune(rune(base[0]))
		return
	}
	printNbrBaseRecursive(nbr, base)
}

func printNbrBaseRecursive(nbr int, base string) {
	if nbr >= StrLen(base) {
		printNbrBaseRecursive(nbr/StrLen(base), base)
	}
	ft.PrintRune(rune(base[nbr%StrLen(base)]))
}