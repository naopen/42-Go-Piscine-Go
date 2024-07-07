/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   enigma.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 23:23:22 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 23:26:31 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Enigma swaps the values pointed to by multiple pointers.
func Enigma(a ***int, b *int, c *******int, d ****int) {
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	***a = tempB
	*b = tempD
	*******c = tempA
	****d = tempC
}
