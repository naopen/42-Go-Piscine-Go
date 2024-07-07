/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   treenode.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/08 03:10:03 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/08 03:11:32 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
