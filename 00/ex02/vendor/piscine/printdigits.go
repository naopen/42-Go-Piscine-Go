/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printdigits.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:47:02 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/10 17:50:09 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintDigits prints the digits from 0 to 9.
func PrintDigits() {
	for i := '0'; i <= '9'; i++ {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}
