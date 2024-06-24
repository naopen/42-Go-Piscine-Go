/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 05:23:04 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/25 05:23:11 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"piscine"
	"ft"
)

func main() {
	ft.PrintRune(piscine.LastRune("Hello!"))
	ft.PrintRune(piscine.LastRune("Salut!"))
	ft.PrintRune(piscine.LastRune("Ola!"))
	ft.PrintRune('\n')
}