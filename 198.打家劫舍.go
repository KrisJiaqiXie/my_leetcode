package my_leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
// dp[i] = dp[i-2] +nums[i]
// init dp[0]=nums[0] dp[1]=max(nums[0],nums[1])
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))

	//init
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	res := dp[1]
	for i := 2; i < len(nums); i++ {
		//如果偷第i间房，则dp[i-2] + nums[i]
		//如果不偷第i间房，则dp[i-1]
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		res = max(res, dp[i])
	}

	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// @lc code=end
