//两个数组的交集
//https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

func intersect(nums1 []int, nums2 []int) []int {
	var ma map[int]int
	var each []int
	if len(nums1) > len(nums2) {
		ma = slice2map(nums2)
		each = nums1
	} else {
		ma = slice2map(nums1)
		each = nums2
	}
	var repeat []int
	//fmt.Println(ma, nums1, nums2)
	for i := 0; i < len(each); i++ {
		if ma[each[i]] > 0 {
			repeat = append(repeat, each[i])
			ma[each[i]] -= 1
		}
	}
	//fmt.Println(repeat)
	return repeat
}

func slice2map(nums []int) map[int]int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	return m
}