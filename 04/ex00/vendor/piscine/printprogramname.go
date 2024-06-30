/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printprogramname.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:41:12 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 01:06:28 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "os"

// PrintProgramName returns the name of the program without the path.
func PrintProgramName() string {
	return os.Args[0][2:]
}