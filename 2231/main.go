package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		temp := i
		sum := i
		for temp != 0 {
			sum += temp % 10
			temp /= 10
		}

		if n == sum {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(0)
}
