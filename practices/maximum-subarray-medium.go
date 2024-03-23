package practices

// https://leetcode.com/problems/maximum-subarray/description/
// medium
// akane
func maxSubArray(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	curSum, maxSum := nums[0], nums[0]

	for _, v := range nums[1:] {
		curSum = max(curSum+v, v)
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}
