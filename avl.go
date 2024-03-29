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
	out += fmt.Sprint(v.value.value, v.sign)
	out += v.right.String()
	out += ")"
	return
}

func (v *vertex) min() *vertex {
	if v.left == nil {
		return v
	} else {
		return v.left.min()
	}
}

func (v *vertex) max() *vertex {
	if v.right == nil {
		return v
	} else {
		return v.right.max()
	}
}

func rotR(v *vertex) (ov *vertex) {
	ov = v.left
	v.left = ov.right
	ov.right = v
	ov.sign = 0
	v.sign = 0
	return
}

func DrotR(v *vertex) (ov *vertex) {
	x, y := v, v.left
	ov = y.right
	x.left = ov.right
	y.right = ov.left
	ov.left = y
	ov.right = x
	switch ov.sign {
	case -1:
		y.sign = 0
		x.sign = 1
	case 0:
		y.sign = 0
		x.sign = 0
	case 1:
		y.sign = -1
		x.sign = 0
	}
	ov.sign = 0
	return
}

func rotL(v *vertex) (ov *vertex) {
	ov = v.right
	v.right = ov.left
	ov.left = v
	ov.sign = 0
	v.sign = 0
	return
}

func DrotL(v *vertex) (ov *vertex) {
	x, y := v, v.right
	ov = y.left
	x.right = ov.left
	y.left = ov.right
	ov.right = y
	ov.left = x
	switch ov.sign {
	case -1:
		y.sign = 1
		x.sign = 0
	case 0:
		y.sign = 0
		x.sign = 0
	case 1:
		y.sign = 0
		x.sign = -1
	}
	ov.sign = 0
	return
}

func (v *vertex) insert(p plate) (ov *vertex, bigger bool) { // chytry insert
	var isBigger bool

	if v == nil {
		ov = &vertex{}
		ov.value = p
		bigger = true
	} else if p.value < v.value.value {
		v.left, isBigger = v.left.insert(p)
		if isBigger {
			switch v.sign {
			case -1:
				switch v.left.sign {
				case -1: // jednoduchá rotace vpravo
					ov = rotR(v)
					bigger = false
				case 1: // dvojitá rotace vpravo url: https://ksp.mff.cuni.cz/kucharky/vyhledavaci-stromy/
					ov = DrotR(v)
					bigger = false
				}
			case 0:
				v.sign = -1
				bigger = true
				ov = v
			case 1:
				v.sign = 0
				bigger = false
				ov = v
			}
		} else {
			ov = v
		}
	} else if p.value > v.value.value {
		v.right, isBigger = v.right.insert(p)
		if isBigger {
			switch v.sign {
			case 1:
				switch v.right.sign {
				case 1: // jednoduchá rotace vlevo
					ov = rotL(v)
					bigger = false
				case -1: // dvojitá rotace vlevo url: https://ksp.mff.cuni.cz/kucharky/vyhledavaci-stromy/
					ov = DrotL(v)
					bigger = false
				}
			case 0:
				v.sign = 1
				bigger = true
				ov = v
			case -1:
				v.sign = 0
				bigger = false
				ov = v
			}
		} else {
			ov = v
		}
	}
	return
}

func (v *vertex) delete(value uint) (ov *vertex, shorter bool) { // chytrý delete
	var isShorter bool

	if value < v.value.value {
		v.left, isShorter = v.left.delete(value)
		if isShorter {
			switch v.sign {
			case -1:
				v.sign = 0
				ov = v
				shorter = true
			case 0:
				v.sign = 1
				ov = v
				shorter = false
			case 1:
				switch v.right.sign {
				case 1: // rotace vlevo
					ov = rotL(v)
					shorter = true
				case 0:
					ov = rotL(v)
					ov.sign = -1
					v.sign = 1
					shorter = false
				case -1: // dvojitá rotace url: https://ksp.mff.cuni.cz/kucharky/vyhledavaci-stromy/
					ov = DrotL(v)
					shorter = true
				}
			}
		} else {
			ov = v
		}
	} else if value > v.value.value {
		v.right, isShorter = v.right.delete(value)
		if isShorter {
			switch v.sign {
			case 1:
				v.sign = 0
				ov = v
				shorter = true
			case 0:
				v.sign = -1
				ov = v
				shorter = false
			case -1:
				switch v.left.sign {
				case -1: // rotace vpravo
					ov = rotR(v)
					shorter = true
				case 0:
					ov = rotR(v)
					ov.sign = 1
					v.sign = -1
					shorter = false
				case 1: // dvojitá rotace url: https://ksp.mff.cuni.cz/kucharky/vyhledavaci-stromy/
					ov = DrotR(v)
					shorter = true
				}
			}
		} else {
			ov = v
		}
	} else if value == v.value.value {
		if (v.right == nil) && (v.left == nil) {
			ov = nil
			shorter = true
		} else if v.right == nil {
			ov = v.left
			shorter = true
		} else if v.left == nil {
			ov = v.right
			shorter = true
		} else {
			tmp := v.right.min()
			v.value = tmp.value
			v.right, isShorter = v.right.delete(tmp.value.value)
			if isShorter {
				switch v.sign {
				case 1:
					v.sign = 0
					ov = v
					shorter = true
				case 0:
					v.sign = -1
					ov = v
					shorter = false
				case -1:
					switch v.left.sign {
					case -1: // rotace vpravo
						ov = rotR(v)
						shorter = true
					case 0:
						ov = rotR(v)
						ov.sign = 1
						v.sign = -1
						shorter = false
					case 1: // dvojitá rotace url: https://ksp.mff.cuni.cz/kucharky/vyhledavaci-stromy/
						ov = DrotR(v)
						shorter = true
					}
				}
			} else {
				ov = v
			}
		}
	}
	return
}
