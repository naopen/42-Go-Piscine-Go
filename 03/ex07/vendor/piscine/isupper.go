/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isupper.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:32:05 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:32:51 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IsUpper checks if a string contains only uppercase letters or is empty.
func IsUpper(s string) bool {
	for _, b := range s {
		if b < 'A' || b > 'Z' {
			return false
		}
	}
	return true
}