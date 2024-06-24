/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   capitalize.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:27:27 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:39:52 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Capitalize capitalizes the first letter of each word and lowercases the rest.
func Capitalize(s string) string {
	result := []byte(s)
	capitalizeNext := true
	for i, b := range result {
		if IsAlpha(string(b)) || IsNumeric(string(b)) {
			if capitalizeNext {
				if IsLower(string(b)) {
					result[i] = b - 32
				}
				capitalizeNext = false
			} else {
				if IsUpper(string(b)) {
					result[i] = b + 32
				}
			}
		} else {
			capitalizeNext = true
		}
	}
	return string(result)
}