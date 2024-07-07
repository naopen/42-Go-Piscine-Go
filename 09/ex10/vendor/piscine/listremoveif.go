/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listremoveif.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:55:59 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:56:04 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListRemoveIf removes all nodes with data equal to data_ref.
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	it := l.Head
	for it.Next != nil {
		if it.Next.Data == data_ref {
			it.Next = it.Next.Next
			if it.Next == nil {
				l.Tail = it
			}
		} else {
			it = it.Next
		}
	}
}
