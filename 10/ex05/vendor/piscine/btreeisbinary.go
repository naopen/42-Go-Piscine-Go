/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreeisbinary.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:28:17 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 03:28:42 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeIsBinary checks if the tree is a binary search tree.
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data >= root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data <= root.Data {
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
