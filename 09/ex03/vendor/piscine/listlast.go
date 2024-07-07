/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listlast.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:26:33 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:26:34 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ListLast returns the last element of the list.
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
