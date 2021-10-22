package my_leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// 递推公式：dp[i]=dp[i-1]+dp[i-2]
// 初始化 dp[1]=1 dp[2]=2
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b, dp := 1, 2, 0
	for i := 3; i <= n; i++ {
		dp = a + b
		a, b = b, dp
	}
	return dp
}

// @lc code=end
