package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [9]int
	sum := 0
	for i := range arr {
		fmt.Scan(&arr[i])
		sum += arr[i]
	}

	sort.Ints(arr[:])

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if sum-arr[i]-arr[j] == 100 {
				for k := 0; k < 9; k++ {
					if k != i && k != j {
						fmt.Println(arr[k])
					}
				}

				return
			}
		}
	}
}
