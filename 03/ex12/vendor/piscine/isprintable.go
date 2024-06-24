/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isprintable.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:33:50 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:34:02 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IsPrintable checks if a string contains only printable characters or is empty.
func IsPrintable(s string) bool {
	for _, b := range s {
		if b < 32 || b > 126 {
			return false
		}
	}
	return true
}