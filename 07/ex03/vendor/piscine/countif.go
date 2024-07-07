/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   countif.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 20:55:09 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 20:55:11 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// CountIf counts the number of elements in a string slice that satisfy a given condition.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}

// IsNumeric checks if a string contains only numeric characters.
func IsNumeric(str string) bool {
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
