package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func minimumBribes(q []int32) {
	fmt.Println(minimumBribesAlgo(q))
}

// Complete the minimumBribes function below.
func minimumBribesAlgo(q []int32) string {

	var initialNextIndexes = []*int{new(int), new(int), new(int)}
	// set next index of the first element
	// * operator is to access the value pointed by pointer
	*initialNextIndexes[2] = 3
	*initialNextIndexes[1] = 2
	*initialNextIndexes[0] = 1

	var possibleNextIndexes = Stacko(initialNextIndexes)

	var totalBribes = 0
	for i, v := range q {
		// fmt.Println(v)
		bribes := possibleNextIndexes.FindAndPull(int(v))
		if bribes < 0 { // index not found
			return "Too chaotic"
		}
		// fmt.Println(bribes)
		totalBribes += bribes
		insertNew := i + 1 + 3
		possibleNextIndexes.Insert(&insertNew)
	}
	var v = strconv.Itoa(totalBribes)
	return v
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

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
