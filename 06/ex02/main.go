/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 05:14:37 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 05:26:32 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"
	"piscine"
)

const (
	usage = "File name missing"
	arg   = "Too many arguments"
)

func main() {
	args := os.Args[1:]
	if piscine.SliceLen(args) == 0 {
		piscine.PrintStr(usage)
	} else if piscine.SliceLen(args) > 1 {
		piscine.PrintStr(arg)
	} else {
		piscine.DisplayFile(args[0])
	}
}
