package basic_sort

func bubblesort(a []int, n int) []int {
	for i:=0; i < n-1; i++ {

		hasSwap := false
		for j:= 0; j < n-i-1; j++ {
			// fmt.Println(a[j], a[j+1])
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				hasSwap = true
			}
		}
		if !hasSwap {
			break
		}
		//fmt.Println(a)
	}
	return a
}