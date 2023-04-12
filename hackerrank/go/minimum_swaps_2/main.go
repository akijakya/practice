package main

import "fmt"

func main() {
	arr := []int32{1, 3, 5, 2, 4, 6, 7}
	fmt.Println(minimumSwaps(arr))
}

func minimumSwaps(arr []int32) (int32, []int32) {
    var count int32
    for i, _ := range arr {
        for (i+1) != int(arr[i]) {
            arr[i], arr[int(arr[i])-1] = arr[int(arr[i])-1], arr[i]
            count++
        }
    }
    return count, arr
}