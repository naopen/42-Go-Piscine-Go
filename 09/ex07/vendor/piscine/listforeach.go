/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listforeach.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:48:34 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:48:35 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListForEach applies a function to each node in the list.
func ListForEach(l *List, f func(*NodeL)) {
	it := l.Head
	for it != nil {
		f(it)
		it = it.Next
	}
}

// Add2_node adds 2 to the data of a node if it is an integer.
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}
