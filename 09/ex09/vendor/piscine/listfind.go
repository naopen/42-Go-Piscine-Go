/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   listfind.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:53:05 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:53:24 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// CompStr compares two interfaces for equality.
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind returns the address of the first node with data equal to ref.
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	it := l.Head
	for it != nil {
		if comp(it.Data, ref) {
			return &it.Data
		}
		it = it.Next
	}
	return nil
}
