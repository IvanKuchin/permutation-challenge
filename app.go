package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
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

func get_permutations_from_number(n int) [][]int {
	var arr []int

	for n > 0 {
		// get the last digit
		d := n % 10
		// append it to the array
		arr = append(arr, d)
		// divide by 10
		n = n / 10
	}

	var result [][]int

	Perm(arr, func(a []int) {
		b := make([]int, len(a))
		copy(b, a)

		result = append(result, b)

	})

	return result
}

func splice_number(n int) []int {
	var arr []int

	n = int(math.Abs(float64(n)))

	for n > 0 {
		// get the last digit
		d := n % 10
		// prepend it to the array
		arr = append([]int{d}, arr...)
		// divide by 10
		n = n / 10
	}

	return arr
}

func contains(arr [][]int, item []int) bool {
	for _, a := range arr {
		if slices.Equal(a, item) {
			return true
		}
	}

	return false
}

func analyze_number(i int) {
	perms := get_permutations_from_number(i)

	for _, perm := range perms {
		subtractor := perm[0]*100 + perm[1]*10 + perm[2]
		subtracted := i - subtractor

		spliced := splice_number(subtracted)

		// fmt.Println(i, subtractor, subtracted, " ", spliced)

		if contains(perms, spliced) {
			fmt.Println("Found a match: ", i, subtractor, subtracted)
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
