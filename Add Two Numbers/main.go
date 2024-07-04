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

}

	}
}



