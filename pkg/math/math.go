// pkg/math/math.go
package math

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pow(x, y int) int {
	if y == 0 {
		return 1
	}

	tmp := Pow(x, y/2)
	if y%2 == 0 {
		return tmp * tmp
	}

	if y > 0 {
		return x * tmp * tmp
	}

	return (tmp * tmp) / x
}
