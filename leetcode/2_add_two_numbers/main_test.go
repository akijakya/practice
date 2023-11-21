package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		description string
		listNode1   *ListNode
		listNode2   *ListNode
		expected    *ListNode
	}{
		{
			description: "equally long inputs",
			listNode1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			listNode2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			description: "equally long inputs with large number in the end",
			listNode1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
			listNode2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
		{
			description: "one input is shorter by 1",
			listNode1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
			listNode2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
		{
			description: "two digits difference in length",
			listNode1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
			listNode2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			res := addTwoNumbers(test.listNode1, test.listNode2)

			for res != nil {
				if res.Val != test.expected.Val {
					t.Errorf("the expected res is %v, received: %v", test.expected, res)
				}
				res = res.Next
				test.expected = test.expected.Next
			}
		})
	}
}
