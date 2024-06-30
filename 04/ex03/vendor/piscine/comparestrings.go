/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   comparestrings.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:50:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:58:35 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// CompareStrings compares two strings lexicographically.
func CompareStrings(a, b string) bool {
	i := 0
	for i < StrLen(a) && i < StrLen(b) {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
		i++
	}
	return StrLen(a) < StrLen(b)
}
