package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	listNode2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(listNode1, listNode2)

	for res != nil {
		println(fmt.Sprintf("%v", res.Val))
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{Val: 0, Next: nil}
	_, _, _ = buildResList(l1, l2, &res)
	return res.Next
}

func buildResList(l1, l2, res *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1.Next == nil && l2.Next != nil {
		res.Next = getNextNode(res, l1, l2)
		return buildResList(&ListNode{Val: 0, Next: nil}, l2.Next, res.Next)
	}

	if l1.Next != nil && l2.Next == nil {
		res.Next = getNextNode(res, l1, l2)
		return buildResList(l1.Next, &ListNode{Val: 0, Next: nil}, res.Next)
	}

	// break recursion
	if l1.Next == nil && l2.Next == nil {
		res.Next = getNextNode(res, l1, l2)

		// deal with if we still have carry
		if res.Next.Val > 9 {
			res.Next.Val = res.Next.Val % 10
			res.Next.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
		return nil, nil, nil
	}

	res.Next = getNextNode(res, l1, l2)
	return buildResList(l1.Next, l2.Next, res.Next)
}

func getNextNode(previous, i1, i2 *ListNode) *ListNode {
	nextNode := &ListNode{}

	// deal with carry
	if previous.Val > 9 {
		previous.Val = previous.Val % 10
		nextNode.Val = i1.Val + i2.Val + 1
	} else {
		nextNode.Val = i1.Val + i2.Val
	}

	return nextNode
}
