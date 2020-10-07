package main

import "fmt"

func main() {
	var T uint
	fmt.Scan(&T)
	for i := uint(0); i < T; i++ {
		var N uint
		var l, a *element
		l = &element{}
		a = l

		fmt.Scan(&N)
		fmt.Scan(&l.value)
		for j := uint(1); j < N; j++ {
			var Ni uint
			fmt.Scan(&Ni)
			a = a.append(Ni)
		}
		fmt.Println(l.String())
	}
}
