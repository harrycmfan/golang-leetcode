package golang_leetcode

func putOnTable(a []int) []int {
	a2 := []int{}
	for len(a) > 0 {
		a2 = append(a2, a[0])
		a = a[1:]
		if len(a) > 1 {
			a = append(a[1:], a[0])
		}
	}
	return a2
}

func putOnHands(a []int) []int {
	a1 := []int{}
	for i := len(a)-1; i >=0; i-- {
		a1 = append(a1, a[i])
	}
	a2 := []int{}
	for len(a1) > 0 {
		if len(a2) > 1 {
			a2 = append([]int{a2[len(a2)-1]}, a2[:len(a2)-1]...)
		}
		a2 = append([]int{a1[0]}, a2...)
		a1 = a1[1:]
	}
	return a2
}
