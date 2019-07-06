package basic_sort

// 把shell排序看成insert排序的进阶
// insert排序是gap为1的shell
func shellsort(a []int, n int) []int {

	for gap := n/2; gap>=1; gap = gap/2 {
		for i:= gap; i <n; i+=gap {
			val := a[i]
			for j:=i-gap; j>=0;j-=gap {
				if a[j] <= val {
					break
				}
				a[j+gap] = a[j]
				a[j] = val
			}
		}
	}

	return a
}
