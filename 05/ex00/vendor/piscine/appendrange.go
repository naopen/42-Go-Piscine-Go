/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   appendrange.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 02:24:42 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 02:24:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// AppendRange returns a slice of integers from min to max (excluding max).
func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}