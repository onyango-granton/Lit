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

func arrToLinkedList(arr []int) *ListNode {
	initNode := &ListNode{}
	for i, ch := range arr {
		if i == 0 {
			initNode.Val = ch
			initNode.Next = nil
		} else {
			currentNode := &ListNode{}
			currentNode.Val = ch
			currentNode.Next = initNode
			initNode = currentNode
		}
	}
	return initNode
}


