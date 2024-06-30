/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   revparams.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:44:43 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:40:07 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// RevParams prints the command line arguments in reverse order.
func RevParams(args []string) {
	for i := SliceLen(args) - 1; i >= 0; i-- {
		PrintStr(args[i] + "\n")
	}
}