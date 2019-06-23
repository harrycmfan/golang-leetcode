package golang_leetcode

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxLength := 1
	var startPos, endPos int
	myMap := map[uint8]int{s[startPos]:startPos}

	for endPos += 1; endPos < len(s); endPos ++{
		if lastIndex, ok := myMap[s[endPos]]; ok {
			if lastIndex >= startPos { // 当window size是1 并且后一个还是跟前一个一样的时候出现
				startPos = lastIndex+1
			}
		}
		myMap[s[endPos]] = endPos
		windowLen := endPos-startPos+1
		if windowLen > maxLength {
			maxLength = windowLen
		}
	}

	// charArray := [128]int{}
	// startPos := 0
	// for endPos := 0; endPos < len(s); endPos++ {
	//     if charArray[s[endPos]] > startPos {
	//         startPos = charArray[s[endPos]]
	//     }
	//     charArray[s[endPos]] = endPos + 1
	//     if endPos-startPos+1 > maxLength {
	//         maxLength = endPos-startPos+1
	//     }
	// }
	return maxLength
}

// 笨办法,从第一个开始遍历,用一个map来记录已经出现的char，继续遍历下一个直到有重复的，记录map的长度
// 继续从第二个开始遍历，每次替换最长的map长度，n2

// 优化办法就是怎么能遍历一次呢
// 使用sliding window，window保持没有重复的substring，不断向右移动
// 优化第一步，使用一个map，保存遍历过的字符，如果没有在map里面，则加到map中
// 如果已经在map里面了，就移除startPos当前的字符，然后startPos前进一格，endPos不变，继续看是不是移除的这个重复的，如果没有，说明startPos还要继续前进移除
/* while (i < n && j < n) {
       // try to extend the range [i, j]
       if (!set.contains(s.charAt(j))){
           set.add(s.charAt(j++));
           ans = Math.max(ans, j - i);
       }
       else {
           set.remove(s.charAt(i++));
       }
   }
*/
// 这个办法是O（2n）
// 更近一步优化，map保存lastIndex，只需要endpos遍历前进，如果字符已经在map中了，那么看一下lastIndex和当前的startPos谁大，如果startPos小，那么移动startPos到当前的endPos,如果startPos还更大，说明之前的这个字符已经在window之外了，就不管了startPos了，直接跟新index即可。

// 最优解，假设都是ascii码，我们使用一个数组，数组默认都是0，说明数字没有出现过
// 从第一个endPos0开始遍历，如果endPos上的值比startPos要大，说明这个数字出现过，更新startPos
// 每次都更新数组pos的值
// 用map记录和用array记录本质上是一样的，都有一个怎么移动startPos的考虑

