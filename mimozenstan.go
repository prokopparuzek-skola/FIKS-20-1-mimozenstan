package main

import "fmt"

func main() {
	var T uint
	fmt.Scan(&T)
	for i := uint(0); i < T; i++ {
		var N, Ni uint
		var l, a *element
		var avl *vertex
		l = &element{}
		a = l

		fmt.Scan(&N)
		fmt.Scan(&Ni)
		l.value = Ni
		avl, _ = avl.insert(plate{Ni, l})
		for j := uint(1); j < N; j++ {
			fmt.Scan(&Ni)
			a = a.append(Ni)
			avl, _ = avl.insert(plate{Ni, a})
			fmt.Println(avl)
		}
		fmt.Println(l)
		fmt.Println(avl)
	}
}
