package Week_08

//191. 位1的个数

func HammingWeight() int {
	var num uint32 = 00000000000000000000000000001011
	return hammingWeight(num)
}

func hammingWeight(num uint32) int {
	//fmt.Println(num)
	var count int
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
