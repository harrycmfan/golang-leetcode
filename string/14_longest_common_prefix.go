package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	res := ""
	for {
		var add byte
		for j := 0; j < len(strs); j++ {
			//fmt.Println(strs, j)
			if len(strs[j]) == i {
				return res
			}
			if j == 0 {
				add = strs[j][i]
			} else {
				if add != strs[j][i] {
					return res
				}
			}
		}
		res += string(add)
		i++
	}
	return res
}

// Vertical scanning