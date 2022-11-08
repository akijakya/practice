package main

import "fmt"

func main() {
	nums1 := []int{2, 3, 4, 5}
	target1 := 9

	println(fmt.Sprintf("result: %v", twoSum(nums1, target1)))
}

func twoSum(nums []int, target int) []int {
	var hashTable map[int]int
	hashTable = make(map[int]int)
	var res []int

	for i, num := range nums {
		if i > 0 {
			if val, ok := hashTable[target-num]; ok {
				res = append(res, i, val)
				break
			}
		}
		hashTable[num] = i
	}

	return res
}
