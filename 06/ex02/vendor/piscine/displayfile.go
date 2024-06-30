/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   displayfile.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 05:19:13 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 05:27:31 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"os"
)

func DisplayFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		PrintStr(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var buffer [1024]byte
	for {
		n, err := file.Read(buffer[:])
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			os.Stdout.Write([]byte{buffer[i]})
		}
	}
}