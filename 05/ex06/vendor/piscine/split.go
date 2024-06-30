/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 03:09:15 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 03:10:11 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Split splits a string by a separator.
func Split(s, sep string) []string {
	var result []string
	wordStart := 0
	for i := 0; i < StrLen(s); i++ {
		if s[i:i+StrLen(sep)] == sep {
			result = append(result, s[wordStart:i])
			wordStart = i + StrLen(sep)
			i += StrLen(sep) - 1
		}
	}
	result = append(result, s[wordStart:])
	return result
}