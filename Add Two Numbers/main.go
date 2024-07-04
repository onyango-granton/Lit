package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

}

func linkedListToArr(l *ListNode) []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func revArr(arr []int) []int {
	res := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}

func longerArr(arr1, arr2 []int) []int {
	if len(arr1) > len(arr2) {
		return arr1
	} else {
		return arr2
	}
}



