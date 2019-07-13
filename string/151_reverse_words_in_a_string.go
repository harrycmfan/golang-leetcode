package string

/*
Given an input string, reverse the string word by word.



Example 1:

Input: "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.
 */

func reverseWords(s string) string {
	res := []string{}
	word := ""
	for i:=0; i < len(s); i++ {
		if s[i] == ' ' && len(word) == 0 {
			continue
		}
		if s[i] == ' ' && len(word) > 0 {
			res = append(res, word)
			word = ""
			continue
		}
		word += string(s[i])
	}
	if len(word) > 0 {
		res = append(res, word)
	}
	if len(res) == 0 {
		return ""
	}
	str := ""
	for i:=len(res)-1; i>=0; i-- {
		str += res[i] + " "
	}
	return str[:len(str)-1]
}

// 更好的方法是reverse string，reverse word, remove space, 3步，顺序可以调换