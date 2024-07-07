/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreeapplyinorder.go                               :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:16:52 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 03:16:54 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeApplyInorder applies a function to each node in the tree in-order.
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
