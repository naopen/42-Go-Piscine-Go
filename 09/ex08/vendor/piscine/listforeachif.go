/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listforeachif.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:50:33 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:51:21 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListForEachIf applies a function to nodes in the list that satisfy a condition.
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if cond(it) {
			f(it)
		}
		it = it.Next
	}
}

// IsPositiveNode checks if a node's data is a positive number.
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

// IsAlNode checks if a node's data is not a number.
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
