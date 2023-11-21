package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		listNode *ListNode
		expected bool
	}{
		{
			listNode: &ListNode{
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
			},
			expected: true,
		},
		{
			listNode: &ListNode{
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
			},
			expected: true,
		},
		{
			listNode: &ListNode{
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
			},
			expected: true,
		},
		{
			listNode: &ListNode{
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
			},
			expected: false,
		},
		{
			listNode: &ListNode{
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
			},
			expected: false,
		},
		{
			listNode: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: true,
		},
		{
			listNode: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			expected: true,
		},
		{
			listNode: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := isPalindrome(test.listNode)

			if res != test.expected {
				t.Errorf("the expected res is %v, received: %v", test.expected, res)
			}
		})
	}
}
