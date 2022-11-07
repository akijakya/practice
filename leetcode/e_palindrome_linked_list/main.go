package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	listNode2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	listNode3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	listNode4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	listNode5 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	listNode6 := &ListNode{
		Val:  1,
		Next: nil,
	}

	listNode7 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	listNode8 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	println(fmt.Sprintf("result: %v", isPalindrome(listNode1)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode2)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode3)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode4)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode5)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode6)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode7)))
	println(fmt.Sprintf("result: %v", isPalindrome(listNode8)))
}

func isPalindrome(head *ListNode) bool {
	var s []int
	_, slice := collectVal(head, s)
	println(fmt.Sprintf("slice: %#v", slice))

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
