package practices

// https://leetcode.com/problems/longest-increasing-subsequence/
// medium
/*
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
// DP method O(n^2)

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	length := make([]int, len(nums))
	length[0] = 1

	for i := 1; i < len(nums); i++ {
		length[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && length[i] < length[j]+1 {
				length[i] = length[j] + 1
			}
		}
	}
	max := 0
	for _, v := range length {
		if v > max {
			max = v
		}
	}

	return max
}
