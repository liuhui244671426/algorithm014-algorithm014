func plusOne(digits []int) []int {
    count := len(digits)
	n := 1
	fmt.Println(digits)
	for i := count; i > 0; i-- {
		if n == 1 && digits[i-1]+n == 10 {
			//满 10 进 1
			digits[i-1] = 0
		} else if n == 1 {
			digits[i-1] = digits[i-1] + n
			n = 0
		}
	}
	if n == 1 {
		digits = append([]int{n}, digits...)
	}
	return digits
}