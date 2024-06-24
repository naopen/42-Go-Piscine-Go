/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   index.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:24:27 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:45:47 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Index returns the index of the first instance of toFind in s, or -1 if not found.
func Index(s string, toFind string) int {
	for i := range s {
		if StrLen(s[i:]) < StrLen(toFind) {
			return -1
		}
		if s[i:i+StrLen(toFind)] == toFind {
			return i
		}
	}
	return -1
}