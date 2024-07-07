/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slicelenint.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:01:16 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:01:52 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// SliceLenInt counts the number of elements in an int slice.
func SliceLenInt(slice []int) int {
	count := 0
	for range slice {
		count++
	}
	return count
}
