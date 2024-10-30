package main

import (
	"fmt"
	"math/rand/v2"
)

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := range(n) {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	var slice []int
	for _=range(100) {
		slice = append(slice, rand.IntN(1000))
	}
	
	
	fmt.Print("Before sorting\n")
	for _, v := range(slice) {
		fmt.Printf("%d, ", v)
	}
	fmt.Print("Before sorting\n")

	slice = insertionSort(slice)

	fmt.Println("\nAfter sorting")
	for _, v := range(slice) {
		fmt.Printf("%d, ", v)
	}
}
