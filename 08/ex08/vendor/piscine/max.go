/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   max.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:04:59 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:05:06 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Max returns the maximum value in a slice of integers.
func Max(a []int) int {
	if SliceLenInt(a) == 0 {
		return 0
	}
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
