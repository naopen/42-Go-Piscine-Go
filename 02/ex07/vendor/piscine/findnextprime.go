/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   findnextprime.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 02:38:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 02:44:32 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine


// FindNextPrime finds the next prime number greater than or equal to a given number.
func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}