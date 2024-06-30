/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strlen.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 12:58:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:57:46 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// StrLen counts the number of runes in a string.
func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
