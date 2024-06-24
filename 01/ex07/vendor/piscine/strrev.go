/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strrev.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:01:06 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 01:54:02 by nkannan          ###   ########.fr       */
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

// StrRev reverses a string.
// 1. Create a rune slice manually and place it in reverse order.
// 2. Convert the rune slice to a string manually.
func StrRev(s string) string {
	length := StrLen(s)
	var reversed string
	for i := length - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
