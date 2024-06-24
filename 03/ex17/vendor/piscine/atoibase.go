/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoibase.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:36:43 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:57:28 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// AtoiBase converts a string to an integer in a given base.
func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := StrLen(base)
	for _, r := range s {
		digit := -1
		for i, b := range base {
			if b == r {
				digit = i
				break
			}
		}
		if digit == -1 {
			return 0
		}
		result = result*baseLen + digit
	}
	return result
}

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