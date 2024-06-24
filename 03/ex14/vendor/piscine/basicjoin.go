/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicjoin.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:34:25 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 06:34:33 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BasicJoin concatenates a slice of strings.
func BasicJoin(elems []string) string {
	result := ""
	for _, s := range elems {
		result += s
	}
	return result
}