package euclid

// Find greatest common divisior between 2 positive integers m, n using Euclid algorithm.
// It must satisfy the precondition of m > 0, n > 0 or else the result is undefined.
func Gcd(m int, n int) (int) {
	r := m%n
	if r == 0 {
		return n
	} 
	return Gcd(n, r)
}
