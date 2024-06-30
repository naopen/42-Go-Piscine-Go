/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   convertbase.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 03:00:49 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 03:04:18 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ConvertBase converts a number from one base to another.
func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := AtoiBase(nbr, baseFrom)
	return PrintNbrBase(decimal, baseTo)
}

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

// PrintNbrBase returns a string representation of an integer in a given base.
func PrintNbrBase(nbr int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}
	if nbr < 0 {
		return "-" + PrintNbrBase(-nbr, base)
	}
	if nbr == 0 {
		return string(base[0])
	}
	return PrintNbrBase(nbr/StrLen(base), base) + string(base[nbr%StrLen(base)]) // ここでStrLen関数を使用してbaseの長さを取得
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
