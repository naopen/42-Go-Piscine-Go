/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreeapplybylevel.go                               :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:30:39 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 04:52:49 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func QueueLen(queue []*TreeNode) int {
	count := 0
	for range queue {
		count++
	}
	return count
}

// BTreeApplyByLevel applies a function to each node in the tree level by level.
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	nodes := []*TreeNode{root}
	for QueueLen(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		f(current.Data)

		if current.Left != nil {
			nodes = append(nodes, current.Left)
		}
		if current.Right != nil {
			nodes = append(nodes, current.Right)
		}
	}
}
