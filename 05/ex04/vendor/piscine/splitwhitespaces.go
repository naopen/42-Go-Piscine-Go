/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   splitwhitespaces.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 02:53:39 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 02:53:41 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// SplitWhiteSpaces splits a string by whitespace characters.
func SplitWhiteSpaces(str string) []string {
	var result []string
	var word string
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}