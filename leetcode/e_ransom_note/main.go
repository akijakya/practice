package main

func main() {
	println(canConstruct("well hoods", "hello world"))
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	res := true

	for _, letter := range magazine {
		if _, ok := magazineMap[letter]; ok {
			magazineMap[letter]++
		} else {
			magazineMap[letter] = 1
		}
	}

	for _, letter := range ransomNote {
		value, ok := magazineMap[letter]

		if !ok {
			res = false
			break
		}

		if value < 1 {
			res = false
			break
		}

		magazineMap[letter]--
	}

	return res
}
