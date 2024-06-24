/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   lastrune.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:23:41 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:45:17 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// LastRune returns the last rune of a string.
func LastRune(s string) rune {
	if StrLen(s) == 0 {
		return 0
	}
	return []rune(s)[StrLen(s)-1]
}