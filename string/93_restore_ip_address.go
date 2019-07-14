package string

import (
	"strconv"
	"strings"
)

/*
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
 */

func restoreIpAddresses(s string) []string {
	res := []string{}
	restore(s, 4, "", &res)
	return res
}

func restore(s string, part int, prefix string, list *[]string) {
	if part == 0 {
		if len(s) == 0 {
			//fmt.Println(prefix[:len(prefix)-1])
			*list = append(*list, prefix[:len(prefix)-1]) //remve last .
		}
		return
	}
	if len(s) < part {
		return
	}
	round := 3 // default 3
	if len(s)-part+1 < 3 { // exmple, len(s) = 4, part = 3, can only cut on 1 or 2
		round = len(s)-part+1
	}
	for i:= 1; i <= round; i++ {
		ip := s[0:i]
		if isValid(ip) {
			restore(s[i:], part-1, prefix+ip+".", list)
		}
	}
}

func isValid(s string) bool {
	if s == "0" {
		return true
	}
	if strings.HasPrefix(s, "0") {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}
	return true
}

// 递归思想
// 1 先写出valid IP 的方法
// 2 然后明确查找有效ip无非是确定4段都是valid，可以简化成确定一段valid，剩下3段，依次类推
//   在4段都valid的情况下，看看还剩下吗？剩下就是不合法，不然就ok了