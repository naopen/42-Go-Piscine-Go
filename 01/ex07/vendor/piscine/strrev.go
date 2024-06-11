/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strrev.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:01:06 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 17:02:12 by nkannan          ###   ########.fr       */
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

	reversedRunes := make([]rune, length)
	index := length - 1
	for _, r := range s {
		reversedRunes[index] = r
		index--
	}

	reversed := ""
	for i := 0; i < length; i++ {
		reversed += string(reversedRunes[i])
	}

	return reversed
}
