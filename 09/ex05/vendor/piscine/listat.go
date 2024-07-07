/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listat.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:36:52 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:41:50 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListAt returns the node at the specified position in the list.
func ListAt(l *NodeL, pos int) *NodeL {
	if pos <= 0 {
		return nil
	}
	count := 0
	it := l
	for it != nil {
		if count == pos {
			return it
		}
		count++
		it = it.Next
	}
	return nil
}
