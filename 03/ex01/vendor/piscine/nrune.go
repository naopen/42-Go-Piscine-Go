/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   nrune.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:22:31 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:44:03 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// NRune returns the nth rune of a string.
func NRune(s string, n int) rune {
	if n <= 0 || n > StrLen(s) {
		return 0
	}
	return []rune(s)[n-1]
}