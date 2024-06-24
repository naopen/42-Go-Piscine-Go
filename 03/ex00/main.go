/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:15:21 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 05:17:56 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"piscine"
	"ft"
)

func main() {
	ft.PrintRune(piscine.FirstRune("Hello!"))
	ft.PrintRune(piscine.FirstRune("Salut!"))
	ft.PrintRune(piscine.FirstRune("Ola!"))
	ft.PrintRune('\n')
}