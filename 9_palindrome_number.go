package golang_leetcode

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 */


func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}
	//originX := x
	reverse := 0
	// for x > 0 {
	//     left := x%10
	//     reverse = reverse * 10 + left
	//     x = x/10
	//     //fmt.Println(x, reverse)
	// }
	// return reverse == originX
	for x > reverse { // trick 在这里，只需要遍历到一半即可
		left := x%10
		reverse = reverse * 10 + left
		x = x/10
		//fmt.Println(x, reverse)
	}
	//fmt.Println(x, reverse)
	return reverse == x || reverse/10 == x
}