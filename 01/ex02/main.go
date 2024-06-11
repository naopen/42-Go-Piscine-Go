/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 12:49:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 12:49:32 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
