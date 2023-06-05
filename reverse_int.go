package reverse_int

func ReverseInt(n int) int {
	x := 0
	for n > 0 {
		x = x*10 + n%10
		n /= 10
	}
	return x
}
