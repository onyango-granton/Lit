package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode)  {
	arr1 := []int{}
	arr2 := []int{}

	revArr1 := []int{}
	revArr2 := []int{}

	// shorterArr := []int{}
	// longerArr := []int{}

	for l1 != nil{
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil{
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}

	for i:=len(arr1)-1; i>=0;i--{
		revArr1 = append(revArr1, arr1[i])
	}

	for i:=len(arr2)-1; i>=0;i--{
		revArr2 = append(revArr2, arr2[i])
	}

	// if len(arr1) > len(arr2){
	// 	shorterArr = arr1
	// } else {
	// 	shorterArr = arr2
	// }

	// if len(arr1) > len(arr2){
	// 	longerArr = arr1
	// } else {
	// 	longerArr = arr2
	// }

	arr1Index := len(arr1)-1
	arr2Index := len(arr2)-1

	resArr := []int{}

	carry := 0

	for !(arr1Index < 0) && !(arr2Index < 0){
		fmt.Println(revArr1[arr1Index],revArr2[arr2Index])
		indexSum := revArr1[arr1Index]+revArr2[arr2Index]
		remainder := indexSum % 10
		indexCarry := indexSum / 10
		resArr = append(resArr, remainder+carry)
		carry = indexCarry
		arr1Index--
		arr2Index--
	}



	fmt.Println(resArr)
}

func main() {
	Anode1 := &ListNode{Val: 7, Next: nil}
	Anode2 := &ListNode{Val: 5, Next: Anode1}
	Anode3 := &ListNode{Val: 3, Next: Anode2}
	Anode4 := &ListNode{Val: 8, Next: Anode3}
	Anode5 := &ListNode{Val: 6, Next: Anode4}
	Anode6 := &ListNode{Val: 5, Next: Anode5}
	Anode7 := &ListNode{Val: 6, Next: Anode6}
	Anode8 := &ListNode{Val: 8, Next: Anode7}
	Anode9 := &ListNode{Val: 0, Next: Anode8}
	//753865680
	//798580876

	Bnode1 := &ListNode{Val: 7, Next: nil}
	Bnode2 := &ListNode{Val: 9, Next: Bnode1}
	Bnode3 := &ListNode{Val: 8, Next: Bnode2}
	Bnode4 := &ListNode{Val: 5, Next: Bnode3}
	Bnode5 := &ListNode{Val: 8, Next: Bnode4}
	Bnode6 := &ListNode{Val: 0, Next: Bnode5}
	Bnode7 := &ListNode{Val: 8, Next: Bnode6}
	Bnode8 := &ListNode{Val: 7, Next: Bnode7}
	Bnode9 := &ListNode{Val: 6, Next: Bnode8}

	Cnode1 := &ListNode{Val: 3, Next: nil}
	Cnode2 := &ListNode{Val: 4, Next: Cnode1}
	Cnode3 := &ListNode{Val: 2, Next: Cnode2}
	// Cnode4 := &ListNode{Val: 1, Next: Cnode3}

	Dnode1 := &ListNode{Val: 4, Next: nil}
	Dnode2 := &ListNode{Val: 6, Next: Dnode1}
	Dnode3 := &ListNode{Val: 5, Next: Dnode2}

	addTwoNumbers(Anode9,Bnode9)
	addTwoNumbers(Cnode3,Dnode3)
}