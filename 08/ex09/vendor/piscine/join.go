/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   join.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:09:26 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:10:55 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Join concatenates strings with a given separator.
func Join(strs []string, sep string) string {
	var result string
	for i, str := range strs {
		result += str
		if i < SliceLenString(strs)-1 {
			result += sep
		}
	}
	return result
}
