package main

import (
	"fmt"
)

func main() {
	var n int
	var a [100]int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
