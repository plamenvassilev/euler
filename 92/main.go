package main

import "fmt"

const max = 10000000

var num, num89, num1 int

// split a number in digitis appended to an array
func splitNumberToDigits(n int) []int {
	var digits []int
	for n > 0 {
		var r = n % 10
		digits = append(digits, r)
		n /= 10
	}
	return digits
}

// Function multiplies each digit by itself and sums the result of all
func squareDigitChains() {
	var nums []int
	var sum int

	for {
		if num == 89 {
			num89 = num89 + 1
			break
		}
		if num == 1 {
			num1 = num1 + 1
			break
		}
		nums = splitNumberToDigits(num)
		for _, n := range nums {
			sum += n * n
			//fmt.Printf("%d * %d = %d \n", n, n, sum)
			num = sum
			//fmt.Println(num)
		}
		sum = 0
	}
}

func main() {
	for i := 1; i <= max; i++ {
		num = i
		squareDigitChains()
	}
	fmt.Println("1s:", num1, "89s", num89)
}

// Problem asks for 89s only but I am returning the 1s as well
// result: 1s: 1418854 89s 8581146

// A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.
//
// For example,
//
// 44 → 32 → 13 → 10 → 1 → 1
// 85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89
//
// Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.
//
// How many starting numbers below ten million will arrive at 89?
