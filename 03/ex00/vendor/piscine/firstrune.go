/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   firstrune.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:16:19 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:43:18 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// FirstRune returns the first rune of a string.
func FirstRune(s string) rune {
	if StrLen(s) == 0 {
		return 0
	}
	return []rune(s)[0]
}
