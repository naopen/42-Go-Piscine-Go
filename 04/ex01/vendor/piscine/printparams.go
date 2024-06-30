/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printparams.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:41:26 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:36:24 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// PrintParams prints the command line arguments.
func PrintParams(args []string) {
	for _, arg := range args {
		PrintStr(arg + "\n")
	}
}