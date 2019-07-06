package basic_sort


func heapsort(a []int, n int) []int {
	for i:= n/2 -1 ;i>=0;i-- {
		a = adjustHeap(a, i, n)
	}


	for i:=n-1; i>=1; i-- {
		tmp := a[i]
		a[i] = a[0]
		a[0] = tmp
		a = adjustHeap(a, 0, i)
	}
	return a
}

func adjustHeap(a []int, parent int, n int) []int{
	child := parent * 2 +1;
	for child < n {
		parentVal := a[parent]
		if child + 1 < n && a[child+1] > a[child] {
			child = child + 1
		}
		if parentVal > a[child] {
			break
		}
		a[parent] = a[child]
		a[child] = parentVal
		parent = child
		child = child * 2 + 1
	}
	return a
}
