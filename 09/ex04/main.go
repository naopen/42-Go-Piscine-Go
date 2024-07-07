/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:31:02 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:31:05 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

type List = piscine.List
type Node = piscine.NodeL

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	piscine.ListPushBack(link, "I")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "something")
	piscine.ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	piscine.ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}
