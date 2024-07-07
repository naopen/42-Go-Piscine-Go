/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreesearchitem.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:24:06 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 03:24:07 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeSearchItem searches for a node with the given data in the tree.
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}
