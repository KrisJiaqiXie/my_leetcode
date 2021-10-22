package my_leetcode

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	var sum int
	for _, value := range nums {
		sum += value
	}
	if sum%2 != 0 { //如果和为奇数，则不可能分成两个相等的数组
		return false
	}

	sum /= 2

	//建立二维数组
	var dp [][]bool
	dp = make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j >= nums[i-1] { //如果容量够用则可放入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else { //如果容量不够用则不拿，维持前一个状态
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}

// @lc code=end
