package my_leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// 递推公式dp[i] = dp[i-1]+dp[i-2]
// 初始化dp[0]=1 dp[1]=1
// 遍历顺序
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// @lc code=end
