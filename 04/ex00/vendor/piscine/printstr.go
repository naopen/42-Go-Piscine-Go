/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printstr.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 12:52:31 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 12:53:01 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

// PrintStr prints one by one the characters of a string on the screen.
func PrintStr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
}
