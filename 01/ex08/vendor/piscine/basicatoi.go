/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicatoi.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:06:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 17:08:20 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BasicAtoi converts a string to an integer.
func BasicAtoi(s string) int {
	result := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result * 10 + int(r - '0')
		} else {
			return 0
		}
	}
	return result
}
