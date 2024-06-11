/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sortintegertable.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:19:15 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 17:19:17 by nkannan          ###   ########.fr       */
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

// SortIntegerTable sorts a slice of integers in ascending order.
func SortIntegerTable(table []int) {
	n := SliceLen(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				// スワップを手動で行う
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}
