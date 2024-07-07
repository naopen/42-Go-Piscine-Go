/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 23:36:30 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 23:46:38 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "ft"

// Door represents a door with a state.
type Door struct {
	state bool
}

const OPEN = true
const CLOSE = false

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

// OpenDoor opens a door.
func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
	return true
}

// CloseDoor closes a door.
func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
	return true
}

// IsDoorOpen checks if a door is open.
func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state == OPEN
}

// IsDoorClose checks if a door is closed.
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
