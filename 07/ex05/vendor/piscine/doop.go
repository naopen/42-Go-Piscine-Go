/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   doop.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:08:19 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:08:20 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"errors"
	"fmt"
)

// Doop performs an arithmetic operation on two integers.
func Doop(arg1 int, operator string, arg2 int) (int, error) {
	switch operator {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0 {
			return 0, errors.New("No division by 0")
		}
		return arg1 / arg2, nil
	case "%":
		if arg2 == 0 {
			return 0, errors.New("No modulo by 0")
		}
		return arg1 % arg2, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}
