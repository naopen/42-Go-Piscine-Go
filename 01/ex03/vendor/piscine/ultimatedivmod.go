/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ultimatedivmod.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 12:51:14 by nkannan           #+#    #+#             */
/*   Updated: 2024/06/11 12:51:23 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// UltimateDivMod divides the int pointed by a by the int pointed by b.
// The result is stored in the int pointed by a.
// The remainder of the division is stored in the int pointed by b.
func UltimateDivMod(a *int, b *int) {
	*a, *b = *a / *b, *a % *b
}
