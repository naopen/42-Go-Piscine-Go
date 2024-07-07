/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   map.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 20:45:52 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 20:51:24 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Map applies a function to each element of an int slice and returns a slice of booleans.
func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}

// IsPrime checks if a number is prime.
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i * i <= nb; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
