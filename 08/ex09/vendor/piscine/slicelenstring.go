/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slicelen.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 00:39:36 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 00:10:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// SliceLen counts the number of elements in a slice.
func SliceLenString(slice []string) int {
    count := 0
    for range slice {
        count++
    }
    return count
}
