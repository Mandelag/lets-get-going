package main

import (
	"testing"
)

func TestEllipsis(t *testing.T) {
	name := []string{"Keenan", "Mandela", "Gebze"}
	// someFun := func (fullName ...string) {
	// 	fullName[0] = "Keenongg"
	// }
	hmm(name...)
	t.Log("hai")

	// https://medium.com/golangspec/variadic-functions-in-go-13c33182b851
	// " Variadic functions are syntactic sugar for functions taking slice as a last parameter. "
	t.Log(name) // ??????????????

	name2a := "Keenan"
	name2b := "Mandela"
	name2c := "Gebze"

	hmm(name2a, name2b, name2c)

	t.Log(name2a, name2b, name2c) // ????

	name2a = "Keenan"
	name2b = "Mandela"
	name2c = "Gebze"

	// hmm(name2a, ([]string{name2b, name2c})...) // ;) not compiling
	//	t.Log(name2a, name2b, name2c) // ???? ;)

	hmm([]string{name2a, name2b, name2c}...) // ;) compiling, but reference are lost

	t.Log(name2a, name2b, name2c) // ???? ;)

}

func hmm(fullName ...string) {
	fullName[0] = "Kinun"
	fullName[1] = "Manderos"
}

// Interesting find in benchmarking prealocating slice size.
// It seems that slice with 0 size, but b.N capacity perform the fastest.
// If the size is already b.N, and the capacity also B.N, its less fast.
func BenchmarkWithoutPrealocation(b *testing.B) {
	// str := make([]int, 0)
	str := make([]string, 0)
	for j := 0; j < b.N; j++ {
		// i, _ := strconv.Atoi("10")
		i := "10" // strconv.Atoi("10")
		str = append(str, i)
	}
}

func BenchmarkWithPrealocation(b *testing.B) {
	// str := make([]int, 0, b.N)
	str := make([]string, 0, b.N)
	for j := 0; j < b.N; j++ {
		// i, _ := strconv.Atoi("10")
		i := "10" // strconv.Atoi("10")
		str = append(str, i)
	}
}
