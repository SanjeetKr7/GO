package main

import "fmt"

func iSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

func main() {
	arr := []int{5, 3, 1, 6, 2}
	fmt.Println(iSort(arr))
}
