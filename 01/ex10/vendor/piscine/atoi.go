/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoi.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:15:48 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 17:15:49 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Atoi converts a string to an integer, handling signs and invalid characters.
func Atoi(s string) int {
	sign := 1
	result := 0
	for i, r := range s {
		if i == 0 && (r == '+' || r == '-') {
			if r == '-' {
				sign = -1
			}
		} else if r >= '0' && r <= '9' {
			result = result * 10 + int(r - '0')
		} else {
			return 0
		}
	}
	return result * sign
}
