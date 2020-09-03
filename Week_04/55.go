package Week_04
//跳跃游戏 
import "fmt"

func CanJump() bool {
	var nums []int = []int{2, 3, 1, 1, 4}
	return canJump(nums)
}

func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println("i,k", i, k)
		if i > k {
			return false
		}
		if k < (i + nums[i]) {
			k = i + nums[i]
		}
	}
	return true
}
