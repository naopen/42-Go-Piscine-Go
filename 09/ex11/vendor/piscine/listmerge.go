/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listmerge.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:57:10 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:57:14 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListMerge merges two lists by appending the second list to the first.
func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		*l1 = *l2
		return
	}

	if l2.Head == nil {
		return
	}

	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
