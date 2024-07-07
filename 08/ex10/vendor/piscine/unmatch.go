/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   unmatch.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:13:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:13:48 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Unmatch finds the element in a slice of ints that does not have a matching pair.
func Unmatch(a []int) int {
	for i := 0; i < SliceLenInt(a); i++ {
		count := 0
		for j := 0; j < SliceLenInt(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count % 2 != 0 {
			return a[i]
		}
	}
	return -1
}
