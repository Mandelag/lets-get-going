package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, Go!")
	// https://docs.oracle.com/javase/7/docs/api/java/lang/String.html
	// https://learning.oreilly.com/library/view/get-programming-with/9781617293092/kindle_split_019.html
	fmt.Println("The string is encoded as UTF-8 variable length string! Not like java which is UTF-16")
	fmt.Println("がんばりましょう。ゴランをべんきょうします")

	// just like python's print() function, golang fmt.Println function accept multiple parameters.
	// The following example will print the string, separated by space.
	fmt.Println("Keenan", "Mandela", "Gebze")
	fmt.Println(len("がんばりましょう"), "bytes")
	var r, size = utf8.DecodeRuneInString("がんばりましょう")
	fmt.Printf("First rune: %c %v bytes", r, size)

	//byte in go is unsigned int8, not like java which is signed byte
	var b byte = 0x20 // space character
	fmt.Printf("%c\n", b)
	decodeSecretMessage("L fdph, L vdz, L frqtxhuhg.")
	fmt.Println("string join with number " + string('1'))
	fmt.Println(10)
	fmt.Println(math.MinInt16)
	var s = strconv.Itoa(32)
	fmt.Println(s)

	hehe()
	var tes = hehe
	fmt.Printf("%[1]T ", tes)
}

func encodeSecretMessage(str string) {

}

func decodeSecretMessage(in string) {
	for _, c := range in {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			fmt.Printf("%c", c-3)
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func hehe() int32 {
	type celcius float64
	var temperature celcius = 10.4
	fmt.Printf("%v\n", temperature)
	// bedakan type aliases dengan type
	return 0
}
