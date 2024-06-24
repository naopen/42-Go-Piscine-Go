/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sqrt.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 02:37:33 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 02:37:37 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Sqrt calculates the square root of a number.
func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}