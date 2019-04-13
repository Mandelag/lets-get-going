package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hai")
	var input = [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}
	fmt.Println(arrayManipulation(5, input))
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	// var result int64

	var movingAverages = make([]float64, n) // make splice of size n
	var ns = make([]int, n)                 // multiplier of average (should start with 1, but..)

	for _, v := range queries {
		// calculate moving average
		var a = v[0] - 1
		var b = v[1] - 1
		var k = v[2]

		for j := a; j <= b; j++ {
			movingAverages[j] = movingAverages[j] + float64(k)/float64(len(queries))
			ns[j] = ns[j] + 1
		}
		// fmt.Println(movingAverages)
	}

	var max float64
	for _, v := range movingAverages {
		if v > max {
			max = v
		}
	}
	// fmt.Println(movingAverages, max)

	result := math.Round(max * float64(len(queries)))
	// return result
	return int64(result)
}
