package my_leetcode

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start

func integerBreak(n int) int {
	//init
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// @lc code=end
