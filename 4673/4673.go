package main

import "fmt"

func generateNumber(n int) int {
	num := n

	for n != 0 {
		num += n % 10
		n /= 10
	}

	return num
}

func main() {
	var isNotSelf [10001]bool

	for i := 1; i <= 10000; i++ {
		num := generateNumber(i)
		if num <= 10000 {
			isNotSelf[num] = true
		}
	}

	for i := 1; i <= 10000; i++ {
		if !isNotSelf[i] {
			fmt.Println(i)
		}
	}
}
