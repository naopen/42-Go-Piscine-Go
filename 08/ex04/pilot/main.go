/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 23:29:02 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 23:29:03 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

// Pilot represents a pilot with name, age, aircraft, and life.
type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft int
}

const AIRCRAFT1 = 1

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}
