/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listreverse.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:43:38 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:43:42 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListReverse reverses the order of elements in the list.
func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
