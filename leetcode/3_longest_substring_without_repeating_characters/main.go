package main

import "fmt"

func main() {
	s := "abcabcbb"
	println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	prevArr := make([]rune, 0)

	for i, c := range []rune(s) {
		arr := make([]rune, 0)
		arr = append(arr, c)

		for j := 0; j < len(prevArr); j++ {
			if prevArr[j] != c {
				arr = append(arr, prevArr[j])
			} else {
				break
			}
		}

		if len(arr) > res {
			res = len(arr)
		}

		prevArr = arr

		println(fmt.Sprintf("%d(%s): %#v", i, string(c), arr))
	}

	return res
}

func lengthOfLongestSubstringOptimal(s string) int {
	// Length of the given input string
	n := len(s)

	// Length of longest substring
	var result int

	// Map to store visited characters along with their index
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			result = max(result, end-start)

			// Remove all characters before the new
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	result = max(result, n-start)

	return result
}
