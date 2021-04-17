package main

import "fmt"

/**
升降字符串 先从小到大 再从大到小排
*/
func main() {
	aa := plusChar("aaaabbbbcccc")
	fmt.Println(aa)
}

func plusChar(s string) string {
	if s == "" {
		return ""
	}
	wordArr := [26]int{}
	for _, v := range s {
		wordArr[v-'a']++
	}
	str := ""
	i := 0
	for len(str) < len(s) {
		if i%2 == 0 {
			for k := 0; k < 26; k++ {
				if wordArr[k] > 0 {
					//从小到大排
					str += string(k + 'a')
					wordArr[k]--
				}

			}

		} else {
			for l := 25; l >= 0; l-- {
				if wordArr[l] > 0 {
					//从小到大排
					str += string(l + 'a')
					wordArr[l]--
				}

			}
		}
		i++
	}

	return str
}
