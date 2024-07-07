/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   comcheck.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 22:52:46 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 22:52:48 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"fmt"
	"os"
)

// ComCheck checks the command line arguments for specific values and prints an alert if any match.
func ComCheck() {
	for _, arg := range os.Args[1:] {
		if arg == "42" || arg == "piscine" || arg == "piscine 42" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
