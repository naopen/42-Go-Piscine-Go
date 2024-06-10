/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isnegative.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:49:45 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/10 17:50:02 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// IsNegative prints 'T' if the input integer is negative, and 'F' otherwise.
func IsNegative(nb int) {
	if nb < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
