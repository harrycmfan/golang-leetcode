package golang_leetcode

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
 */

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	return findMaxHeight(height)
}

func findMaxHeight(a []int) int {
	n := len(a)
	maxHeightIndex := 0
	maxHeight := a[0]

	for i:=0;i<n;i++ {
		if a[i] > maxHeight {
			maxHeightIndex = i
			maxHeight = a[i]
		}
	}

	sum := 0
	leftMax := 0
	for i:= 0; i < maxHeightIndex; i++ {
		if a[i] > leftMax {
			leftMax = a[i]
		}
		if a[i] < leftMax {
			sum+=leftMax-a[i]
		}
	}

	rightMax := 0
	for i:= n-1; i > maxHeightIndex; i-- {
		if a[i] > rightMax {
			rightMax = a[i]
		}
		if a[i] < rightMax {
			sum+=rightMax-a[i]
		}
	}
	return sum
}

func oneLoop(a []int) int {
	n := len(a)
	leftMax := 0
	rightMax := 0

	i := 0
	j := n-1
	sum := 0

	for i <= j {
		if a[i] <= a[j] {
			if a[i] > leftMax {
				leftMax = a[i]
			} else if a[i] < leftMax {
				sum += leftMax - a[i]
			}
			i++
		} else {
			if a[j] > rightMax {
				rightMax = a[j]
			} else if a[j] < rightMax {
				sum += rightMax - a[j]
			}
			j--
		}

	}
	return sum
}

func three_loop(a []int) int {
	n := len(a)

	leftMaxArray := make([]int, n)
	rightMaxArray := make([]int, n)

	leftMaxArray[0] = 0
	leftMaxArray[1] = a[0]
	for i:=2; i < n; i++ {
		if leftMaxArray[i-1] > a[i-1] {
			leftMaxArray[i] = leftMaxArray[i-1]
		} else {
			leftMaxArray[i] = a[i-1]
		}
	}
	//fmt.Println(leftMaxArray)

	rightMaxArray[n-1] = 0
	rightMaxArray[n-2] = a[n-1]
	for i:=n-3; i >= 0; i-- {
		if rightMaxArray[i+1] > a[i+1] {
			rightMaxArray[i] = rightMaxArray[i+1]
		} else {
			rightMaxArray[i] = a[i+1]
		}
	}

	sum := 0
	for i := 1; i < n-1; i++ {
		minHeight := leftMaxArray[i]
		if rightMaxArray[i] < leftMaxArray[i] {
			minHeight = rightMaxArray[i]
		}
		if minHeight > a[i] {
			sum+=minHeight-a[i]
		}
	}
	return sum
}
