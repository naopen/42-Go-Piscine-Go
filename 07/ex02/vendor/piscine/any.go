/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   any.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 20:52:42 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 20:52:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Any checks if any element in a string slice satisfies a given condition.
func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
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
