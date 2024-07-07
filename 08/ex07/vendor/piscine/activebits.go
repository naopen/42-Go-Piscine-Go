/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   activebits.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:01:24 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:01:29 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ActiveBits counts the number of active bits in an integer.
func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		if n & 1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
