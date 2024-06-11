/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   divmod.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 12:49:41 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 12:49:43 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// DivMod divides a by b and stores the result in the int pointed by div.
// The remainder of the division is stored in the int pointed by mod.
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
