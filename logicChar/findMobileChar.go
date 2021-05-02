package main

import (
	"fmt"
	"strings"
)

/**
输入一个电话号码 返回数字所代表的所有的字母组合
*/
func main() {
	dig := "235"
	res := []string{}

	str := ""
	digMap := make(map[string]string)
	digMap["2"] = "abc"
	digMap["3"] = "def"
	digMap["4"] = "ghi"
	digMap["5"] = "jkl"
	digMap["6"] = "mno"
	digMap["7"] = "pqrs"
	digMap["8"] = "tuv"
	digMap["9"] = "wxyz"

	findAllChar(dig, 0, str, digMap, &res)
	fmt.Println(res)
}

func findAllChar(digital string, index int, str string, digMap map[string]string, res *[]string) {
	if len(digital) == 0 {
		return
	}
	if index == len(digital) {
		*res = append(*res, str)
		return
	}

	chaNum := string(digital[index])
	chas := digMap[chaNum]

	for _, v := range chas {
		str += string(v)
		findAllChar(digital, index+1, str, digMap, res)
		l1 := strings.Count(str, "") - 1
		str = str[:l1-1]
	}

}
