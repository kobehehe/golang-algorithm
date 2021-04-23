package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{30, 12, 13, 3, 20, 8, 5}
	arr1 := quickSort(arr)
	fmt.Println(arr1)
}

func quickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	flag := arr[0]
	var left, right []int
	for i := 1; i < n; i++ {
		if arr[i] < flag {
			left = append(left, arr[i])
		} else if arr[i] > flag {
			right = append(right, arr[i])
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		left = quickSort(left)
		wg.Done()
	}()

	go func() {
		right = quickSort(right)
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
