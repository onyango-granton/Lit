package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resArr := []uint{}

	res := listToNum(l1) + listToNum(l2)
	fmt.Println(listToNum(l1))
	fmt.Println(res)
	count := 0

	for res != 0 {
		remainder := res % 10
		resArr = append(resArr, remainder)
		res /= 10
	}

	resNode := &ListNode{}

	for i := len(resArr) - 1; i >= 0; i-- {
		count++
		if count == 1 {
			resNode.Val = int(resArr[i])
			resNode.Next = nil
		} else {
			currentNode := &ListNode{Val: int(resArr[i]), Next: resNode}
			resNode = currentNode
		}

	}

	return resNode
}

// func revuint(num uuint) uuint {
// 	res := 0
// 	for num != 0{
// 		remainder := num % 10
// 		res = res * 10 + remainder
// 		num /= 10
// 	}
// 	return res
// }

func listToNum(l *ListNode) uint {
	var res uint
	numArr := []int{}
	for l != nil {
		numArr = append(numArr, l.Val)
		l = l.Next
	}
	for i := len(numArr) - 1; i >= 0; i-- {
		res = res*10 + uint(numArr[i])
	}
	fmt.Println(res)
	return res
}

func printNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main() {
	// Anode1 := &ListNode{Val: 7, Next: nil}
	// Anode2 := &ListNode{Val: 5, Next: Anode1}
	// Anode3 := &ListNode{Val: 3, Next: Anode2}
	// Anode4 := &ListNode{Val: 8, Next: Anode3}
	// Anode5 := &ListNode{Val: 6, Next: Anode4}
	// Anode6 := &ListNode{Val: 5, Next: Anode5}
	// Anode7 := &ListNode{Val: 6, Next: Anode6}
	// Anode8 := &ListNode{Val: 8, Next: Anode7}
	// Anode9 := &ListNode{Val: 0, Next: Anode8}
	// //753865680
	// //798580876

	// Bnode1 := &ListNode{Val: 7, Next: nil}
	// Bnode2 := &ListNode{Val: 9, Next: Bnode1}
	// Bnode3 := &ListNode{Val: 8, Next: Bnode2}
	// Bnode4 := &ListNode{Val: 5, Next: Bnode3}
	// Bnode5 := &ListNode{Val: 8, Next: Bnode4}
	// Bnode6 := &ListNode{Val: 0, Next: Bnode5}
	// Bnode7 := &ListNode{Val: 8, Next: Bnode6}
	// Bnode8 := &ListNode{Val: 7, Next: Bnode7}
	// Bnode9 := &ListNode{Val: 6, Next: Bnode8}

	Cnode30 := &ListNode{Val: 1, Next: nil}
	Cnode29 := &ListNode{Val: 0, Next: Cnode30}
	Cnode28 := &ListNode{Val: 0, Next: Cnode29}
	Cnode27 := &ListNode{Val: 0, Next: Cnode28}
	Cnode26 := &ListNode{Val: 0, Next: Cnode27}
	Cnode25 := &ListNode{Val: 0, Next: Cnode26}
	Cnode24 := &ListNode{Val: 0, Next: Cnode25}
	Cnode23 := &ListNode{Val: 0, Next: Cnode24}
	Cnode22 := &ListNode{Val: 0, Next: Cnode23}
	Cnode21 := &ListNode{Val: 0, Next: Cnode22}
	Cnode20 := &ListNode{Val: 0, Next: Cnode21}
	Cnode19 := &ListNode{Val: 0, Next: Cnode20}
	Cnode18 := &ListNode{Val: 0, Next: Cnode19}
	Cnode17 := &ListNode{Val: 0, Next: Cnode18}
	Cnode16 := &ListNode{Val: 0, Next: Cnode17}
	Cnode15 := &ListNode{Val: 0, Next: Cnode16}
	Cnode14 := &ListNode{Val: 0, Next: Cnode15}
	Cnode13 := &ListNode{Val: 0, Next: Cnode14}
	Cnode12 := &ListNode{Val: 0, Next: Cnode13}
	Cnode11 := &ListNode{Val: 0, Next: Cnode12}
	Cnode10 := &ListNode{Val: 0, Next: Cnode11}
	Cnode9 := &ListNode{Val: 0, Next: Cnode10}
	Cnode8 := &ListNode{Val: 0, Next: Cnode9}
	Cnode7 := &ListNode{Val: 0, Next: Cnode8}
	Cnode6 := &ListNode{Val: 0, Next: Cnode7}
	Cnode5 := &ListNode{Val: 0, Next: Cnode6}
	Cnode4 := &ListNode{Val: 0, Next: Cnode5}
	Cnode3 := &ListNode{Val: 0, Next: Cnode4}
	Cnode2 := &ListNode{Val: 0, Next: Cnode3}
	Cnode1 := &ListNode{Val: 1, Next: Cnode2}

	Dnode3 := &ListNode{Val: 4, Next: nil}
	Dnode2 := &ListNode{Val: 6, Next: Dnode3}
	Dnode1 := &ListNode{Val: 5, Next: Dnode2}

	// printNode(Anode3)
	// printNode(Bnode3)
	// var newNo uint = 1000000000000000000000000000001
	// num := new(big.Int)
	// num.
	printNode(addTwoNumbers(Cnode1,Dnode1))
	// fmt.Println(newNo)
}
