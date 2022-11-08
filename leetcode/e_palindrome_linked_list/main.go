package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	println(fmt.Sprintf("result: %v", isPalindrome(listNode)))
}

func isPalindrome(head *ListNode) bool {
	var s []int
	_, slice := collectVal(head, s)

	res := true

	for i, e := range slice {
		if e != slice[len(slice)-i-1] {
			res = false
			break
		}

		if i+1 > len(slice)/2 {
			break
		}
	}
	return res
}

func collectVal(n *ListNode, s []int) (*ListNode, []int) {
	// break recursion
	if n.Next == nil {
		return nil, append(s, n.Val)
	}

	return collectVal(n.Next, append(s, n.Val))
}
