/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   compact.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 23:55:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 23:55:38 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Compact removes empty strings from a string slice.
func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str != "" {
			(*ptr)[count] = str
			count++
		}
	}
	*ptr = (*ptr)[:count]
	return count
}
