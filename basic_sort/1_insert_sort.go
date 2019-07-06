package basic_sort


// 把第一个元素看成有序的，后面的往前插入
func insertsort(a []int, n int) []int {

	for i:= 1; i<n; i++ {
		val := a[i]
		for j := i-1; j >= 0; j-- {
			//fmt.Println(a[j], val)
			if a[j] <= val {
				break
			}
			a[j+1] = a[j]
			a[j] = val
		}
		//fmt.Println(a)
	}
	return a
}