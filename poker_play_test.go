package golang_leetcode

import (
	"fmt"
	"testing"
)

func TestPoker(t *testing.T) {
	fmt.Println(putOnTable([]int{1,2,3,4,5,6,7,8,9}))
	fmt.Println(putOnHands([]int{1,3,5,7,9,4,8,6,2}))
}
