package main

import "fmt"

func main() {
	var T uint
	fmt.Scan(&T)
	for i := uint(0); i < T; i++ {
		var N uint
		var pads []uint

		fmt.Scan(&N)
		pads = make([]uint, N)
		for j := uint(0); j < N; j++ {
			fmt.Scan(&pads[j])
		}
		fmt.Println(N, pads)
	}
}
