/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:35:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:36:36 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "how are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))
}
