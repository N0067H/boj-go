package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	m -= 45

	if m < 0 {
		h--
		m = 60 + m
	}

	if h < 0 {
		h = 23
	}

	fmt.Printf("%d %d", h, m)
}
