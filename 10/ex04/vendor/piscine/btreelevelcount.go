/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreelevelcount.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:26:11 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 03:26:15 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeLevelCount returns the number of levels in the tree.
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
