package Week_11

//860. 柠檬水找零

func LemonadeChange() bool {
	var bills []int
	bills = []int{
		5, 5, 5, 5, 20, 20, 5, 5, 5, 5,
	}
	bills = []int{
		5, 5, 5, 10, 5, 5, 10, 20, 20, 20,
	}
	bills = []int{
		5, 5, 5, 10, 10, 20, 10, 20, 5, 5,
	}
	return lemonadeChange(bills)
}

func lemonadeChange(bills []int) bool {
	var five, ten int = 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		}
		if bill == 10 {
			ten++
		}
		
		if bill > 5 {
			if five == 0 {
				return false
			}
			five--
			if bill > 10 {
				if ten >= 1 {
					ten--
				} else if five >= 2 {
					five -= 2
				} else {
					return false
				}
			}
		}
	}
	return true
}
