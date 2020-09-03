package Week_04

//柠檬水找零
func LemonadeChange() bool {
	var bills []int = []int{10, 10}
	return lemonadeChange(bills)
}

func lemonadeChange(bills []int) bool {
	k_5, k_10, k_20 := 0, 0, 0
	for i := 0; i < len(bills); i++ {
		//fmt.Println(bills[i])
		if bills[i]-5 == 0 {
			k_5++
		} else if bills[i]-5 == 5 {
			if k_5 < 1 {
				return false
			}
			k_5--
			k_10++
		} else {
			if k_5 > 0 && k_10 > 0 {
				k_5--
				k_10--
				k_20++
			} else if k_5 >= 3 {
				k_5 -= 3
				k_20++
			} else {
				return false
			}
		}
	}
	return true
}
