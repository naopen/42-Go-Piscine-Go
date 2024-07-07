/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sortlistinsert.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:58:57 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:59:00 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// NodeI represents a node in a linked list of integers.
type NodeI struct {
	Data int
	Next *NodeI
}

// SortListInsert inserts a new node with data_ref into the sorted list.
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil {
		return newNode
	}

	if data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	return l
}
