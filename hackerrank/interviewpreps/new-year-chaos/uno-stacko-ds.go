package main

/*
*  Container for possible index value
 */
type stacko []*int

/**
* Insert from the tail, return index placed
 */
func (p stacko) insert(value *int) (index int) {
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

/**
* Remove index, reorder, return value removed.
 */
func (p stacko) pull(index int) (value *int) {
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

/**
* Find value, return its index, delete, and reorder.
 */
func (p stacko) findAndPull(value int) (index int) {
	result := -1
	for i := 0; i < len(p); i++ {
		if p[i] != nil && *p[i] == value {
			p.pull(i)
			result = i
			break
		}
	}
	return result
}
