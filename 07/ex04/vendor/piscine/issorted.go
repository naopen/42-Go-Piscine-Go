/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   issorted.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:00:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:01:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// IsSorted checks if an int slice is sorted according to a given comparison function.
func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < SliceLenInt(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
