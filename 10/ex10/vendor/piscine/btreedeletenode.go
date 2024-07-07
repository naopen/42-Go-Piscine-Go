/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   btreedeletenode.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 05:00:06 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 05:00:09 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// BTreeDeleteNode deletes a node from the tree.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil {
		return transplant(root, node, node.Right)
	} else if node.Right == nil {
		return transplant(root, node, node.Left)
	} else {
		successor := BTreeMin(node.Right)
		if successor.Parent != node {
			root = transplant(root, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		root = transplant(root, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
		return root
	}
}

func transplant(root, u, v *TreeNode) *TreeNode {
	if u.Parent == nil {
		return v
	}
	if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
	return root
}
