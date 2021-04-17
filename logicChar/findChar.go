package main

import "fmt"

/**
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。
例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
*/

func main() {
	strArr := []string{"bella", "label", "roller"}
	findChar2(strArr)
}

/**
高级算法 通过文字的ASK码结合数组来搞
*/
func findChar2(A []string) []string {
	wordArr := [26]int{}
	var res []string
	for _, v := range A[0] {
		wordArr[v-'a']++
	}

	for i := 1; i < len(A); i++ {
		wordArr1 := [26]int{}
		for _, v := range A[i] {
			wordArr1[v-'a']++
		}

		for j := 0; j < len(wordArr); j++ {
			if wordArr1[j] < wordArr[j] {
				wordArr[j] = wordArr1[j]
			}
		}
	}

	for i := 0; i < len(wordArr); i++ {
		for j := 0; j < wordArr[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}

/**
基础算法 用map的解法。
*/
func findChar(A []string) []string {
	var res []string
	if len(A) == 0 {
		return res
	}
	baseMap := map[rune]int{}
	compareMap := map[rune]int{}

	for _, v := range A[0] {
		baseMap[v]++
	}

	for i := 1; i < len(A); i++ {
		compareMap = map[rune]int{}
		for _, v := range A[i] {
			compareMap[v]++
		}

		for w, n := range baseMap {
			value, ok := compareMap[w]
			if ok {
				if value < n {
					baseMap[w] = value
				}
			} else {
				delete(baseMap, w)
			}
		}
	}
	for w, n := range baseMap {
		for i := 0; i < n; i++ {
			res = append(res, string(w))
		}
	}
	fmt.Println(res)

	return res
}
