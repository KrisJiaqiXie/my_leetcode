package my_leetcode

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
// 递推公式 dp[i] = min(dp[i-1],dp[i-2])+cost[i]
// 初始化dp[0] = cost[0] dp[1]=cost[1]
// 末尾不好处理，那么就最后两位取个最小值
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, n)

	//init
	dp[0] = cost[0]
	dp[1] = cost[1]

	//loop
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

// @lc code=end
