/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ultimatepointone.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 05:46:16 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 05:46:18 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// UltimatePointOne takes a pointer to a pointer to a pointer to an int
// and sets the value to 1.
func UltimatePointOne(n ***int) {
	***n = 1
}
