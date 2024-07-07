/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listsize.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:13:27 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:15:11 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListSize returns the number of elements in the list.
func ListSize(l *List) int {
	size := 0
	it := l.Head
	for it != nil {
		size++
		it = it.Next
	}
	return size
}
