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

func (v *vertex) min() *vertex {
	if v == nil {
		return v
	} else {
		return v.left.min()
	}
}

func (v *vertex) max() *vertex {
	if v == nil {
		return v
	} else {
		return v.right.max()
	}
}

func (v *vertex) insert(p plate) *vertex { // hloupy insert
	if v == nil {
		v = &vertex{}
		v.value = p
	} else if p.value < v.value.value {
		v.left = v.left.insert(p)
	} else if p.value > v.value.value {
		v.right = v.right.insert(p)
	}
	return v
}

func (v *vertex) delete(value uint) *vertex { // hloupý delete
	if value == v.value.value {
		return nil
	} else if value < v.value.value {
		v.left = v.left.delete(value)
	} else if value > v.value.value {
		v.right = v.right.delete(value)
	} else if value == v.value.value {
		if v.right == nil && v.left == nil {
			return nil
		} else if v.right == nil {
			return v.left
		} else if v.left == nil {
			return v.right
		} else {
			tmp := v.right.min()
			v.value = tmp.value
			v.right = v.right.delete(tmp.value.value)
		}
	}
	return v
}
