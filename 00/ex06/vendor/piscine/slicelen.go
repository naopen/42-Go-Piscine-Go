/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slicelen.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 00:11:46 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 00:23:59 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// SliceLen counts the number of elements in a slice of integers.
func SliceLen(slice []int) int {
	count := 0
	for range slice {
		count++
	}
	return count
}
