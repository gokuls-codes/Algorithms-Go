package main

import (
	"fmt"
	"math/rand/v2"
)

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, len(left) + len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}
	return res
}

func main() {

	var slice []int
	for range 20 {
		slice = append(slice, rand.IntN(1000))
	}

	fmt.Print("Before sorting\n")
	for _, v := range slice {
		fmt.Printf("%d, ", v)
	}

	slice = mergeSort(slice)

	fmt.Println("\nAfter sorting")
	for _, v := range slice {
		fmt.Printf("%d, ", v)
	}
}