package Week_12

//941. 有效的山脉数组

func ValidMountainArray() bool {
	var A []int
	A = []int{0, 3, 2, 1}
	return validMountainArray(A)
}

func validMountainArray(A []int) bool {

	if len(A) < 3 {
		return false
	}
	i, j := 0, len(A)-1
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	}
	for j-1 >= 0 && A[j-1] > A[j] {
		j--
	}
	if i != 0 && i == j && j != len(A)-1 {
		return true
	}
	return false
}
