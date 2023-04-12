package main

import "fmt"

func main() {
	// a := [][]int32{  
	// 	{2, 6, 8},
	// 	{3, 5, 7},
	// 	{1, 8, 1},
	// 	{5, 9, 15},
	// }
	// b := [][]int32{
	// 	{1, 2, 100},
	// 	{2, 5, 100},
	// 	{3, 4, 100},
	// }
	c := [][]int32{
		{29, 40, 787},
		{9, 26, 219},
		{21, 31, 214},
		{8, 22, 719},
		{15, 23, 102},
		{11, 24, 83},
		{14, 22, 321},
		{5, 22, 300},
		{11, 30, 832},
		{5, 25, 29},
		{16, 24, 577},
		{3, 10, 905},
		{15, 22, 335},
		{29, 35, 254},
		{9, 20, 20},
		{33, 34, 351},
		{30, 38, 564},
		{11, 31, 969},
		{3, 32, 11},
		{29, 35, 267},
		{4, 24, 531},
		{1, 38, 892},
		{12, 18, 825},
		{25, 32, 99},
		{3, 39, 107},
		{12, 37, 131},
		{3, 26, 640},
		{8, 39, 483},
		{8, 11, 194},
		{12, 37, 502},
	}
	// fmt.Println(arrayManipulation(10, a))
	// fmt.Println(arrayManipulation(5, b))
	fmt.Println(arrayManipulation(40, c))
}

func arrayManipulation(n int32, queries [][]int32) int64 {
    var arr []int64
    arr = make([]int64, n, n)
    for i := 0; i < len(queries); i++ {
        arr[queries[i][0] - 1] = arr[queries[i][0] - 1] + int64(queries[i][2])
		if int(queries[i][1]) < len(arr) {
			arr[queries[i][1]] = arr[queries[i][1]] - int64(queries[i][2])
		}
		fmt.Println(arr)
    }
	var acc int64
	var res int64
    for j := 0; j < len(arr); j++ { 
		acc = acc + arr[j]
        if acc > res { 
            res = acc
        } 
    }
    return res
}