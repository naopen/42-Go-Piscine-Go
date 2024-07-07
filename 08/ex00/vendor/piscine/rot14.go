/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   rot14.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:34:19 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:34:23 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Rot14 applies the ROT14 cipher to a string.
func Rot14(s string) string {
	result := []rune(s)
	for i, r := range result {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			result[i] = rotate14(r)
		}
	}
	return string(result)
}

func rotate14(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return (((r - 'a') + 14) % 26) + 'a'
	} else if r >= 'A' && r <= 'Z' {
		return (((r - 'A') + 14) % 26) + 'A'
	}
	return r
}
