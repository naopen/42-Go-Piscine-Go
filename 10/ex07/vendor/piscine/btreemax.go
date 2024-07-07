/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreemax.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 04:54:55 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 04:54:58 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeMax returns the node with the maximum value in the tree.
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)
}
