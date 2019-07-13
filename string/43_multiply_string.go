package string

import "fmt"

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	resultArray := make([]int, l1+l2)

	for i := l1-1; i >= 0; i-- {
		n1 := num1[i]-48 // ascii
		for j:= l2-1; j >= 0; j-- {
			n2 := num2[j]-48 // ascii
			mul := int(n1 * n2)
			currPos := i+j+1
			advancePos := currPos-1
			sum := mul + resultArray[currPos] // 这个地方非常的trick，这里目的是把advancePos大于10的情况加起来
			//fmt.Println(mul, resultArray[currPos], resultArray[advancePos])

			resultArray[advancePos] += sum / 10 // advancePos可能大于10
			resultArray[currPos] = (sum) % 10

			// 不需要考虑进位
			//fmt.Println(mul, resultArray[currPos], resultArray[advancePos])

			//             adding := mul/10

			//             resultArray[currPos] += mul%10
			//             if resultArray[currPos] >= 10 {
			//                 resultArray[currPo                             s] -= 10
			//                 adding++
			//             }

			//             resultArray[advancePos] += adding
			// if resultArray[advancePos] >= 10 {
			//     resultArray[advancePos] -= 10
			//     resultArray
			// }
		}
	}
	//fmt.Println(resultArray)
	startZero := true
	str := ""
	for i:= 0; i<len(resultArray); i++ {
		if resultArray[i] == 0 && startZero {
			continue
		}
		startZero = false
		str = str + fmt.Sprintf("%d", resultArray[i])
	}
	return str
}