package basicSort

func mergesort(a []int, n int) []int {
	tmp := make([]int, n)
	return msort(a, 0, n-1, tmp)
}

func msort(a []int, start, end int, tmp []int) []int {
	if start < end {
		mid := (end + start) / 2
		a = msort(a, start, mid, tmp)
		a = msort(a, mid+1, end, tmp)
		return merge(a, start, mid, end, tmp)
	}
	return a
}

func merge(a []int, start, mid, end int, tmp []int) []int {
	i := start
	j := mid + 1
	k := 0
	for i < mid+1 && j <= end {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
			k++
		} else {
			tmp[k] = a[j]
			j++
			k++
		}
	}
	for i < mid+1 {
		tmp[k] = a[i]
		i++
		k++
	}
	for j <= end {
		tmp[k] = a[j]
		j++
		k++
	}
	for i:= 0; i < end-start+1; i++ {
		a[i + start] = tmp[i]
	}
	return a
}