package Week_13

//50. Pow(x, n)

func MyPow() float64 {
	var x float64
	var n int
	x = 2.0
	n = -10
	return myPow(x, n)
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quick(x, n)
	}
	return 1.0 / quick(x, -n)
}

func quick(x float64, n int) float64 {
	var result float64 = 1.0
	xx := x
	for n > 0 {
		if n%2 == 1 {
			result *= xx
		}
		xx *= xx
		n /= 2
	}
	return result
}
