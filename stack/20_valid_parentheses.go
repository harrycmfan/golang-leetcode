package stack

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
 */

func isValid(s string) bool {
	var stack []byte
	for i:=0; i < len(s); i++ {
		switch string(s[i]){
		case `(`, `[`, `{`:
			stack = append(stack, s[i])
		case `)`,`]`,`}`:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1:]
			if !canClose(string(pop), string(s[i])) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func canClose(first, last string) bool {
	if first == `(` && last == `)` {
		return true
	}
	if first == `[` && last == `]` {
		return true
	}
	if first == `{` && last == `}` {
		return true
	}
	return false
}

// push to stack for left, pop for right