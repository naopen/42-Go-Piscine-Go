/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   foreach.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/06 23:43:43 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 20:41:49 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// ForEach applies a function to each element of an int slice.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
	ft.PrintRune('\n')
}

// PrintNbr prints an integer.
func PrintNbr(n int) {
	if n < 0 {
		ft.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	ft.PrintRune(rune(n % 10 + '0'))
}
