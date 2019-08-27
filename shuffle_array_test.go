package golang_leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8,9,10}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println(a)

	a = []int{1,2,3,4,5,6,7,8,9,10}

	rand.Seed(time.Now().UnixNano())
	for i:= 1; i < len(a); i++ {
		swapPos := rand.Intn(i+1)
		//fmt.Println(i, swapPos)
		a[i], a[swapPos] = a[swapPos], a[i]
	}
	fmt.Println(a)
}
