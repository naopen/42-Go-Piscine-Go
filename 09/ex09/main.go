/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 01:52:29 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 01:52:35 by nkannan          ###   ########.fr       */
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
	piscine.ListPushBack(link, "hello1")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")

	found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}
