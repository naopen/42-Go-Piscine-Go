/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   makerange.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 02:27:43 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 02:29:35 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// MakeRange returns a slice of integers from min to max (excluding max).
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	result := make([]int, max-min)
	for i := range result {
		result[i] = min + i
	}
	return result
}
