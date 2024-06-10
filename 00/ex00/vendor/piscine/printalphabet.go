/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printalphabet.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:11:45 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/10 17:22:41 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintAlphabet prints the lowercase alphabet from 'a' to 'z'.
func PrintAlphabet() {
	for r := 'a'; r <= 'z'; r++ {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}