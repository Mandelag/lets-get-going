package main

import "fmt"

// https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

/**
 * container for possible index value
 */
type pic []*int

/**
 * Insert from the tail, return index placed
 */
func (p pic) insert(value *int) (index int) {
	result := -1
	for i := len(p) - 1; i >= 0; i-- { // insert from right (tail) until it finds non-nil value
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
func (p pic) remove(index int) (value *int) {
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
func (p pic) findAndRemove(value int) (index int) {
	result := -1
	for i := 0; i < len(p); i++ {
		if p[i] != nil && *p[i] == value {
			p.remove(i)
			result = i
			break
		}
	}
	return result
}

func main() {
	var a, b, c int = 1, 2, 3
	// fmt.Println(a, b, c)
	var initialNextIndexes = []*int{&a, &b, &c}
	// fmt.Println(initialNextIndexes)
	var chaoticQueue = []int{2, 1, 5, 4, 3}
	var possibleNextIndexes = pic(initialNextIndexes)

	// fmt.Println(possibleNextIndexes)
	// for i := 0; i < len(possibleNextIndexes); i++ {
	// 	if possibleNextIndexes[i] != nil {
	// 		fmt.Print(" ", *possibleNextIndexes[i])
	// 	} else {
	// 		fmt.Print(" ", "-")
	// 	}
	// }
	// fmt.Println()
	var totalBribes = 0
	for i, v := range chaoticQueue {
		// fmt.Println(v)
		bribes := possibleNextIndexes.findAndRemove(v)
		if bribes < 0 { // index not found
			fmt.Println("Too chaotic!")
			break
		}
		fmt.Println(bribes)
		totalBribes += bribes
		insertNew := i + 1 + 3
		possibleNextIndexes.insert(&insertNew)
	}
	fmt.Println("Total:", totalBribes)

}
