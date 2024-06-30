/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   point.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/01 04:48:47 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/01 04:48:49 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Point represents a point in a two-dimensional plane.
type Point struct {
	X int
	Y int
}

// SetPoint sets the coordinates of a point.
func SetPoint(ptr *Point) {
	ptr.X = 42
	ptr.Y = 21
}