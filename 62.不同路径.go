package my_leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
// 递推公式dp[i][j] = dp[i][j-1]+dp[i-1][j]
// 初始化dp[i][0] = 1 dp[0][j] = 1
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	//创建二维数组
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//init
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
