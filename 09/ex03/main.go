/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:25:38 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:25:42 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")

	fmt.Println(piscine.ListLast(link))
	fmt.Println(piscine.ListLast(link2))
}
