/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   toupper.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:26:39 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 05:26:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ToUpper converts all lowercase letters in a string to uppercase.
func ToUpper(s string) string {
	result := []rune(s)
	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = r - ('a' - 'A')
		}
	}
	return string(result)
}