/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listpushback.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:56:13 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:56:15 by nkannan          ###   ########.fr       */
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

// ListPushBack inserts a new node at the end of the list.
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
