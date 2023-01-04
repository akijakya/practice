package main

func main() {
	println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	var count int
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		count++
	}
	return count
}
