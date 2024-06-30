/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printprogramname.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:41:12 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:32:39 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "os"

// PrintProgramName returns the name of the program without the path.
func PrintProgramName() string {
	fullPath := os.Args[0]
	programName := ""

	// Find the last slash in the fullPath
	for i := range fullPath {
		if fullPath[i] == '/' {
			programName = ""
		} else {
			programName += string(fullPath[i])
		}
	}

	return programName
}