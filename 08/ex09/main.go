/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 00:09:13 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:09:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
