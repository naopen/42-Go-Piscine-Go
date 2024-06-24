/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   islower.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:33:11 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:33:38 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IsLower checks if a string contains only lowercase letters or is empty.
func IsLower(s string) bool {
	for _, b := range s {
		if b < 'a' || b > 'z' {
			return false
		}
	}
	return true
}