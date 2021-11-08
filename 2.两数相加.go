package my_leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start

//Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l1 != nil && l2 == nil {
		return l1
	}

	ans := &ListNode{Val: 0}
	res := ans
	epochSum := 0
	carry := 0

	for {
		// handle carry
		epochSum = l1.Val + l2.Val + carry
		if epochSum > 9 {
			carry = 1
			ans.Val = epochSum - 10
		} else {
			carry = 0
			ans.Val = epochSum
		}

		if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
				ans.Next = &ListNode{Val: 1}
			}
			return res
		}

		ans.Next = &ListNode{Val: 0}
		ans = ans.Next

		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}

		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

	}
}

// @lc code=end
