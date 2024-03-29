package practices

import "slices"

// https://leetcode.com/problems/top-k-frequent-elements/
// medium
/*
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]
*/
func topKFrequent(nums []int, k int) []int {
	occurrences := make(map[int]int)

	for _, v := range nums {
		occurrences[v]++
	}

	sorted := make([][]int, len(nums)+1)

	for num, occ := range occurrences {
		sorted[occ] = append(sorted[occ], num)
	}

	res := make([]int, 0, k)
	for i := len(sorted) - 1; i >= 0; i-- {
		slices.Sort(sorted[i])
		if len(res) < k {
			res = append(res, sorted[i]...)
		} else {
			break
		}
	}

	return res
}
