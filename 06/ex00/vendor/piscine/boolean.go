/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   boolean.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 04:39:58 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 04:44:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IsEven checks if an integer is even.
func IsEven(nbr int) bool {
	if nbr % 2 == 0 {
		return true
	}
	return false
}
