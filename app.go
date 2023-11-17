package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

type number struct {
	digits []int
}

func (num *number) make(param int) {
	n := int(math.Abs(float64(param)))

	for n > 0 {
		// get the last digit
		d := n % 10
		// append it to the array
		num.digits = append([]int{d}, num.digits...)
		// divide by 10
		n = n / 10
	}
}

func (num *number) swap(i, j int) {
	num.digits[i], num.digits[j] = num.digits[j], num.digits[i]
}

// Check if a number is in an array
func (num number) in(array []number) bool {
	for _, elem := range array {
		if slices.Equal(elem.digits, num.digits) {
			return true
		}
	}
	return false
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

	num := new(number)
	num.make(n)

	perm(num.digits, func(a []int) {
		b := make([]int, len(a))
		copy(b, a)

		result = append(result, b)

	}, 0)

	return result
}

func contains(arr [][]int, item int) bool {
	num := new(number)
	num.make(item)

	for _, a := range arr {
		if slices.Equal(a, num.digits) {
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
	for i := 954; i < 955; i++ {
		analyze_number(i)
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}
