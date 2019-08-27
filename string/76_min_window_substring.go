package string

/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

 */

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	a := map[byte]int{}
	for i := 0; i < len(t); i++ {
		a[t[i]]++
	}
	left := 0
	for ; left < len(s); left++ {
		if a[s[left]] > 0 {
			break
		}
	}
	if len(t) > len(s)-left {
		return ""
	}
	right := left
	requireCount := len(t)
	minWindow := len(s) + 1
	minLeft := 0
	minRight := 0
	for ; right < len(s); right++ {
		//fmt.Println(left, right, a, requireCount, minLeft, minRight)
		count, ok := a[s[right]]
		if ok {
			a[s[right]]--
			if count > 0 {
				requireCount--
				if requireCount == 0 {
					if right - left + 1 < minWindow {
						minWindow = right - left + 1
						minLeft = left
						minRight = right
					}
					for ; left < right; {
						//fmt.Println(left, right, a, requireCount, minLeft, minRight)
						_, ok := a[s[left]]
						if ok {
							a[s[left]]++
							if a[s[left]] > 0 {
								if right - left + 1 < minWindow {
									minWindow = right - left + 1
									minLeft = left
									minRight = right
								}
								requireCount++
								for left < right {
									left++
									_, ok := a[s[left]]
									if ok {
										break
									}
								}
								break
							} else {
								left++
							}
						} else {
							left++
						}
					}
				}
			}
		}
	}
	if minWindow < len(s) + 1 {
		return s[minLeft:minRight+1]
	}
	return ""
}
