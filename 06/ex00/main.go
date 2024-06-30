/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 03:56:24 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 04:38:53 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"
	"piscine"
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

func main() {
	if piscine.IsEven(piscine.SliceLen(os.Args[1:])) {
		piscine.PrintStr(EvenMsg)
	} else {
		piscine.PrintStr(OddMsg)
	}
}