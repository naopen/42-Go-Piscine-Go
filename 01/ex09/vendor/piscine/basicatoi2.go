/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicatoi2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:13:09 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 17:13:11 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BasicAtoi2 converts a string to an integer, handling non-digit characters.
func BasicAtoi2(s string) int {
	result := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return result
}
