package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func minimumBribes(q []int32) {
	fmt.Println(minimumBribesAlgo(q))
}

// Complete the minimumBribes function below.
func minimumBribesAlgo(q []int32) string {
	
	var initialNextIndexes = []*int{new(int), new(int),new(int)}
	// set next index of the first element
	// * operator is to access the value pointed by pointer
	*initialNextIndexes[2] = 3
	*initialNextIndexes[1] = 2
	*initialNextIndexes[0] = 1

	var possibleNextIndexes = pic(initialNextIndexes)

	var totalBribes = 0
	for i, v := range q {
		// fmt.Println(v)
		bribes := possibleNextIndexes.findAndRemove(int(v))
		if bribes < 0 { // index not found
			return "Too chaotic"
		}
		// fmt.Println(bribes)
		totalBribes += bribes
		insertNew := i + 1 + 3
		possibleNextIndexes.insert(&insertNew)
	}
	var v = strconv.Itoa(totalBribes)
	return v
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

// 	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
// 	checkError(err)
// 	t := int32(tTemp)

// 	for tItr := 0; tItr < int(t); tItr++ {
// 		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
// 		checkError(err)
// 		n := int32(nTemp)

// 		qTemp := strings.Split(readLine(reader), " ")

// 		var q []int32

// 		for i := 0; i < int(n); i++ {
// 			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
// 			checkError(err)
// 			qItem := int32(qItemTemp)
// 			q = append(q, qItem)
// 		}

// 		minimumBribes(q)
// 	}
// }

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * container for possible index value
 */
type pic []*int

/**
 * Insert from the tail, return index placed
 */
func (p pic) insert(value *int) (index int) {
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
