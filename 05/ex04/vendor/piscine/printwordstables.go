/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printwordstables.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 02:55:51 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 02:56:39 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// PrintWordsTables prints each element of a string slice on a separate line.
func PrintWordsTables(a []string) {
	for _, word := range a {
		PrintStr(word + "\n")
	}
}