package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	for _, v := range a {
		if v < x {
			fmt.Print(v, " ")
		}
	}
}
