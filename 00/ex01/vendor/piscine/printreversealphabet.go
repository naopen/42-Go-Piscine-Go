/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printreversealphabet.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:40:03 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/10 17:50:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintReverseAlphabet prints the lowercase alphabet in reverse order from 'z' to 'a'.
func PrintReverseAlphabet() {
	for r := 'z'; r >= 'a'; r-- {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
