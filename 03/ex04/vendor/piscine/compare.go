/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   compare.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:25:27 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:46:00 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Compare compares two strings lexicographically.
func Compare(a, b string) int {
	for i := 0; i < StrLen(a) && i < StrLen(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if StrLen(a) < StrLen(b) {
		return -1
	} else if StrLen(a) > StrLen(b) {
		return 1
	}
	return 0
}