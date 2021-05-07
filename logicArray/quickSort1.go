package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{10, 8, 13, 30, 5, 9, 13}
	targetArr := quickSort1(arr)
	fmt.Println(targetArr)
}

func quickSort1(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}
	flag := arr[0]
	var left, right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < flag {
			left = append(left, arr[i])
		} else if arr[i] > flag {
			right = append(right, arr[i])
		}
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		left = quickSort1(left)
		wg.Done()
	}()

	go func() {
		right = quickSort1(right)
		wg.Done()
	}()
	wg.Wait()
	list := []int{flag}
	if len(right) > 0 {
		list = append(list, right...)
	}

	if len(left) > 0 {
		list = append(left, list...)
	}
	return list
}
