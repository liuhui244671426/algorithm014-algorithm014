package Week_12

//349. 两个数组的交集

func Intersection() []int {
	var nums1 []int = []int{4, 9, 5}
	var nums2 []int = []int{9, 4, 9, 8, 4}

	return intersection(nums1, nums2)
}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	var m map[int]int = make(map[int]int)
	var hash map[int]int = make(map[int]int)
	var nums []int

	if len(nums1) < len(nums2) {
		nums = nums1

		for _, v := range nums2 {
			hash[v]++
		}
	} else {
		nums = nums2
		for _, v := range nums1 {
			hash[v]++
		}
	}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v]++
			if hash[v] > 0 {
				res = append(res, v)
			}

		}
	}
	//fmt.Println(m, hash)
	return res
}
