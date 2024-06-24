/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   join.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 06:34:41 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 07:06:29 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i != 0 {
			result += sep
		}
		result += s
	}
	return result
}