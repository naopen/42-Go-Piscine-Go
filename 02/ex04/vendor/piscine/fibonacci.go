/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fibonacci.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 02:37:11 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 02:37:12 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Fibonacci calculates the Fibonacci number at a given index recursively.
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index <= 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}