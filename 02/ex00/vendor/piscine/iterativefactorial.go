/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   iterativefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 02:34:45 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 02:34:55 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IterativeFactorial calculates the factorial of a number iteratively.
func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}