/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listsort.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:58:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:58:17 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// NodeI represents a node in a linked list of integers.
type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the list in ascending order.
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	head := l
	for head.Next != nil {
		current := head.Next
		for current != nil {
			if head.Data > current.Data {
				head.Data, current.Data = current.Data, head.Data
			}
			current = current.Next
		}
		head = head.Next
	}

	return l
}
