/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:32:08 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:32:26 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"piscine"
	"ft"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
