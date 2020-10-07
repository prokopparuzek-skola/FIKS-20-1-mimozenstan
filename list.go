package main

import "fmt"

type element struct {
	left  *element
	right *element
	value uint
}

func (l *element) String() (out string) {
	if l == nil {
		return
	}
	out += fmt.Sprint(l.value)
	out += ","
	out += l.right.String()
	if l.left == nil {
		out = out[:len(out)-1]
	}
	return
}

func (e *element) append(v uint) *element {
	e.right = &element{e, nil, v}
	return e.right
}

func (e *element) delete() {
	if e.right != nil {
		e.right.left = e.left
	}
	if e.left != nil {
		e.left.right = e.right
	}
}
