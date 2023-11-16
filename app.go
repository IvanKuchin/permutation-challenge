package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

func int_to_array(n int) []int {
	var arr []int

	n = int(math.Abs(float64(n)))

	for n > 0 {
		// get the last digit
		d := n % 10
		// append it to the array
		arr = append([]int{d}, arr...)
		// divide by 10
		n = n / 10
	}

	return arr
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Get all permutations of a number
func get_permutations_from_number(n int) [][]int {
	var result [][]int

	arr := int_to_array(n)

	perm(arr, func(a []int) {
		b := make([]int, len(a))
		copy(b, a)

		result = append(result, b)

	}, 0)

	return result
}

func contains(arr [][]int, item int) bool {
	item_arr := int_to_array(item)

	for _, a := range arr {
		if slices.Equal(a, item_arr) {
			return true
		}
	}

	return false
}

func analyze_number(money_in_the_pocket int) {
	perms := get_permutations_from_number(money_in_the_pocket)

	for _, perm := range perms {
		sushi := perm[0]*100 + perm[1]*10 + perm[2]
		change := money_in_the_pocket - sushi

		if contains(perms, change) {
			fmt.Println("Found a match: ", money_in_the_pocket, sushi, change)
		}
	}
}

func main() {
	start := time.Now()
	for i := 100; i < 999; i++ {
		analyze_number(i)
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}
