package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arrl1 := linkedListToArr(l1)
	arrl2 := linkedListToArr(l2)

	revArrl1 := revArr(arrl1)
	revArrl2 := revArr(arrl2)

	longerArr := longerArr(revArrl1, revArrl2)

	for l1 != nil {
		arrl1 = append(arrl1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		arrl2 = append(arrl2, l2.Val)
		l2 = l2.Next
	}
	lastIndexArrl1 := len(revArrl1) - 1
	lastIndexArrl2 := len(revArrl2) - 1

	carry := 0
	sumArr := []int{}
	for lastIndexArrl1 >= 0 && lastIndexArrl2 >= 0 {
		sum := revArrl1[lastIndexArrl1] + revArrl2[lastIndexArrl2] + carry
		remainder := sum % 10
		carry = sum / 10
		sumArr = append(sumArr, remainder)
		lastIndexArrl1--
		lastIndexArrl2--
	}

	revSumArr := revArr(sumArr)

	lastIndex := len(longerArr) - len(revSumArr)

	longerArr = append(longerArr[:lastIndex], revSumArr...)

	for i := lastIndex - 1; i >= 0; i-- {
		sum := longerArr[i] + carry
		remainder := sum % 10
		carry = sum / 10
		longerArr[i] = remainder
	}

	if carry != 0 {
		newArr := []int{carry}
		newArr = append(newArr, longerArr...)
		longerArr = newArr
	}

	return arrToLinkedList(longerArr)
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


