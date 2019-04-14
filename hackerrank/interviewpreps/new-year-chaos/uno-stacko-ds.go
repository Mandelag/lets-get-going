package main

// "Uno stacko" like data structure.
// You can only place new item from the top.
// You can find and pull the element based on their value.
// Everytime you pull a value, all value on top will drop.
//
// Many of the methods are using value receiver, which is fine because Stacko type
// are based on splice.
// 
// *int type is used rather than int to support nil value of empty stack.
type Stacko []*int

// Place new value on top of the stack.
func (p Stacko) Insert(value *int) (index int) {
	result := -1
	for i := len(p) - 1; i >= 0; i-- { // insert from right (tail) until it finds..
		if p[i] != nil {
			break
		}
		result = i
	}
	p[result] = value
	return result
}

// Pull (or take out) a value from an index index out of stack.
func (p Stacko) Pull(index int) (value *int) {
	result := p[index]
	for i := index; i < len(p); i++ { // remove and shift all to left
		if i+1 >= len(p) {
			p[i] = nil
		} else {
			p[i] = p[i+1]
		}
	}
	return result
}

// Find a value and pull the first found value out of the stack.
func (p Stacko) FindAndPull(value int) (index int) {
	result := -1
	for i := 0; i < len(p); i++ {
		if p[i] != nil && *p[i] == value {
			p.Pull(i)
			result = i
			break
		}
	}
	return result
}
