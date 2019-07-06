package basic_sort

func quicksort(a []int, n int) []int {
	return qsort(a, 0, n-1)
}

// 选择第一个数作为pivot，如果右边的比他小，那么右边的和他换
// 如果左边的比他大，那么左边的和他换，目的就是小的放到左边，大的放到右边
func qsort(a []int, start, end int) []int {
	if end <= start {
		return a
	}
	pivot := a[start]
	i := start
	j := end
	for i < j {
		for i < j && a[j] >= pivot {
			j--
		}
		a[i] = a[j]
		for i < j && a[i] <= pivot {
			i++
		}
		a[j] = a[i]
	}
	a[i] = pivot
	a = qsort(a, start, i-1)
	a = qsort(a, i+1, end)
	return a
}
