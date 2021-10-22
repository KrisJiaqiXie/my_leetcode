package my_leetcode

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {
	a := []int{2, 3, 4, 5, 6}
	res := canPartition(a)
	fmt.Println(res)
}
