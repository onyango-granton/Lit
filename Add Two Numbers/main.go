package main

import "fmt"

type Node struct{
	data int
	nextNode *Node
}

type LinkedList struct{
	headNode *Node
}


func (l *LinkedList) addNode(num int){
	newNode := &Node{data:num}
	currentNode := l.headNode
	l.headNode = newNode
	l.headNode.nextNode = currentNode
}

func (l *LinkedList) printNode() []int{
	res :=
	currentNode := l.headNode
	for currentNode != nil{
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func main() {
	linkList := &LinkedList{}

	linkList.addNode(47)
	linkList.addNode(43)
	linkList.addNode(50)

	linkList.printNode()
}