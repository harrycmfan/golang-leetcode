package basic_sort

// select the min then put at the remaining first pos
func selectsort(a []int, n int) []int {
	for i:=0; i < n-1; i++ {
		min := a[i]
		pos := i
		for j:=i+1;j<n;j++ {
			if a[j] < min {
				min = a[j]
				pos = j
			}
		}
		a[pos] = a[i]
		a[i] = min
	}
	return a
}