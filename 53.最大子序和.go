package my_leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for index, value := range nums {
		if index == 0 {
			continue
		} else {
			dp[index] = max(dp[index-1]+value, value)
			res = max(res, dp[index])
		}
	}
	return res
}
// @lc code=end
