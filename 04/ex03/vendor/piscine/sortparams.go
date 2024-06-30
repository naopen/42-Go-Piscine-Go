/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sortparams.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/30 23:47:01 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 00:58:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// SortParams sorts the command line arguments in ASCII order and prints them.
func SortParams(args []string) {
    length := 0
    for range args {
        length++
    }

    // Implement a simple bubble sort algorithm
    for i := 0; i < length-1; i++ {
        for j := 0; j < length-i-1; j++ {
            if CompareStrings(args[j+1], args[j]) {
                // Swap args[j] and args[j+1]
                args[j], args[j+1] = args[j+1], args[j]
            }
        }
    }

    for _, arg := range args {
        PrintStr(arg + "\n")
    }
}
