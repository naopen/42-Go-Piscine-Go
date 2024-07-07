/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   median.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nkannan <nkannan@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/07 21:43:21 by nkannan           #+#    #+#             */
/*   Updated: 2024/07/07 21:43:26 by nkannan          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// Median returns the median of five integers.
func Median(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	bubbleSort(nums)
	return nums[2]
}

// bubbleSort sorts an array of integers using the bubble sort algorithm.
func bubbleSort(nums []int) {
	n := SliceLenInt(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
