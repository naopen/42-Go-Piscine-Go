/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listclear.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:33:38 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:33:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListClear removes all nodes from the list.
func ListClear(l *List) {
	l.Head = nil
	l.Tail = nil
}
