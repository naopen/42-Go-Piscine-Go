/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:43:24 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:43:25 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)

	piscine.ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
