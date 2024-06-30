package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	// num1 := 0
	// num2 := 0

	// arr1 := []int{}
	// arr2 := []int{}

	resArr := []int{}
	
	// for l1 != nil{
	// 	arr1 = append(arr1, l1.Val)
	// 	// num1 = num1 * 10 + l1.Val
	// 	l1 = l1.Next
	// }
	// for l2 != nil{
	// 	arr2 = append(arr2, l2.Val)
	// 	// num2 = num2 * 10 + l2.Val
	// 	l2 = l2.Next
	// }

	// for i:= len(arr1)-1; i >= 0;i--{
	// 	num1 = num1 * 10 + arr1[i]
	// }

	// for i:= len(arr2)-1; i >= 0;i--{
	// 	num2 = num2 * 10 + arr2[i]
	// }

	

	// res := num1 + num2
	res := listToNum(l1) 
	fmt.Println(res)
	// fmt.Println(num1,num2)
	// fmt.Println(res)
	count := 0

	for res != 0{
		remainder := res % 10
		resArr = append(resArr, remainder)
		res /= 10
	}

	resNode := &ListNode{}

	// fmt.Println(res)

	for i:=len(resArr)-1; i >= 0; i--{
		count++
		if count == 1{
			resNode.Val = resArr[i]
			resNode.Next = nil
		} else {
			currentNode := &ListNode{Val: resArr[i],Next: resNode}
			resNode = currentNode
		}

	}

	// for res != 0{
	// 	remainder := res % 10
	// 	count ++
	// 	if count == 1{
	// 		resNode.Val = remainder
	// 		resNode.Next = nil
	// 		fmt.Println(resNode)
	// 	} else {
	// 		// prevNode := resNode
	// 		// resNode.Val = remainder
	// 		// resNode.Next = prevNode
	// 		// fmt.Println(resNode)
	// 		currentNode := &ListNode{Val: remainder, Next: resNode}
	// 		resNode = currentNode
	// 	}
	// 	res /= 10
	// }

	

	return resNode
}


func listToNum (l *ListNode) int {
	res := 0
	numArr := []int {}
	for l != nil{
		numArr = append(numArr, l.Val)
		l = l.Next
	}
	for i:=len(numArr)-1;i>=0;i--{
		res = res * 10 + numArr[i]
	}
	return res
}


func revInt(num int) int {
	res := 0
	for num != 0{
		remainder := num % 10
		res = res * 10 + remainder
		num /= 10
	}
	return res
}

 func printNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
 }


 func main(){
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

	// printNode(Anode3)
	// printNode(Bnode3)
	printNode(addTwoNumbers(Anode9,Bnode9))
	// addTwoNumbers(Anode3,Bnode3)

	fmt.Println(revInt(345))
	fmt.Println(listToNum(Anode9))
 }
