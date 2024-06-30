/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   concatparams.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 02:50:36 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 02:50:37 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// ConcatParams concatenates a slice of strings with newlines.
func ConcatParams(args []string) string {
	var result string
	for i, arg := range args {
		result += arg
		if i < SliceLen(args)-1 {
			result += "\n"
		}
	}
	return result
}