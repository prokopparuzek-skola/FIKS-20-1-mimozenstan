package main

import "fmt"

type plate struct {
	value    uint
	position *element
}

type vertex struct {
	left  *vertex
	right *vertex
	sign  int
	value plate
}

func (v *vertex) String() (out string) {
	if v == nil {
		return
	}
	out += "("
	out += v.left.String()
	out += fmt.Sprint(v.value.value)
	out += v.right.String()
	out += ")"
	return
}
