package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

var solutions_found = 0

type number struct {
	digits []int
}

func make_number(param int) number {
	var num number
	num.digits = nil
	n := int(math.Abs(float64(param)))

	for n > 0 {
		// get the last digit
		d := n % 10
		// append it to the array
		num.digits = append([]int{d}, num.digits...)
		// divide by 10
		n = n / 10
	}

	return num
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
func perm(a number, f func(number), i int) {
	if i > len(a.digits) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a.digits); j++ {
		a.swap(i, j)
		perm(a, f, i+1)
		a.swap(i, j)
	}
}

// Get all permutations of a number
func get_permutations_from_number(n int) []number {
	var result []number
	num := make_number(n)

	perm(num, func(a number) {
		b := a
		b.digits = nil
		b.digits = append(b.digits, a.digits...)

		result = append(result, b)

	}, 0)

	return result
}

func (num number) get_int() int {
	result := 0
	for _, digit := range num.digits {
		result = result*10 + digit
	}
	return result
}

func same_len(a, b, c int) bool {
	return len(make_number(a).digits) == len(make_number(b).digits) && len(make_number(a).digits) == len(make_number(c).digits)
}

func analyze_number(money_in_the_pocket int) {
	perms := get_permutations_from_number(money_in_the_pocket)

	for _, perm := range perms {
		sushi := perm.get_int()
		change := make_number(money_in_the_pocket - sushi)

		if change.in(perms) && (money_in_the_pocket-sushi) > 0 && same_len(money_in_the_pocket, sushi, money_in_the_pocket-sushi) {
			solutions_found++
			fmt.Println("Found a match: ", money_in_the_pocket, sushi, money_in_the_pocket-sushi)
		}
	}
}

func main() {
	start := time.Now()
	for i := 100; i < 1000; i++ {
		analyze_number(i)
	}
	fmt.Println("Time elapsed: ", time.Since(start))
	fmt.Println("Solutions found: ", solutions_found)
}
