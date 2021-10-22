package my_leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	my_map := make(map[int]int)
	res := []int{0, 0}
	for i, v := range nums {
		m_v, ok := my_map[target-v]
		if ok {
			res[0] = m_v
			res[1] = i
			break
		} else {
			my_map[v] = i
		}
	}
	return res
}

// @lc code=end
