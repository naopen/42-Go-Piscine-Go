/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:18:04 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 05:22:24 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"piscine"
	"ft"
)

func main() {
	ft.PrintRune(piscine.NRune("Hello!", 3))
	ft.PrintRune(piscine.NRune("Salut!", 2))
	ft.PrintRune(piscine.NRune("Bye!", -1))
	ft.PrintRune(piscine.NRune("Bye!", 5))
	ft.PrintRune(piscine.NRune("Ola!", 4))
	ft.PrintRune('\n')
}