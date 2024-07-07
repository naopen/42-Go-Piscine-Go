/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:03:24 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:08:06 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if piscine.SliceLen(os.Args) != 4 {
		return
	}

	arg1 := piscine.Atoi(os.Args[1])
	operator := os.Args[2]
	arg2 := piscine.Atoi(os.Args[3])

	result, err := piscine.Doop(arg1, operator, arg2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
