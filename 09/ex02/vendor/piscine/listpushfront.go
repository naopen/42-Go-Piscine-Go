/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listpushfront.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:57:51 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:57:53 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// NodeL represents a node in a linked list.
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents a linked list.
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushFront inserts a new node at the beginning of the list.
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
