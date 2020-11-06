package Week_08

import (
	"fmt"
	"math"
)

func Sort() []int {
	var nums []int
	nums = []int{5, 2, 34, 1, 2, 4, 54, 2, 3, 4, 5, 65, 6}
	fmt.Println("input : ", nums)
	//r := quickSort(nums, 0, len(nums)-1)
	r := radixSort(nums)
	return r
}

/*
冒泡排序

算法描述
1.比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3.针对所有的元素重复以上的步骤，除了最后一个 (len-i)；
4.重复步骤1~3，直到排序完成
*/
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

/*
选择排序

算法描述
n个记录的直接选择排序可经过n-1趟直接选择排序得到有序结果。具体算法描述如下：

1.初始状态：无序区为R[1..n]，有序区为空；
2.第i趟排序(i=1,2,3…n-1)开始时，当前有序区和无序区分别为R[1..i-1]和R(i..n）。该趟排序从当前无序区中-选出关键字最小的记录 R[k]，将它与无序区的第1个记录R交换，使R[1..i]和R[i+1..n)分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区；
3.n-1趟结束，数组有序化了。
*/
func selectionSort(nums []int) []int {
	var minIndex int
	for i := 0; i < len(nums)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

/*
插入排序


算法描述
一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：

1.从第一个元素开始，该元素可以认为已经被排序；
2.取出下一个元素，在已经排序的元素序列中从后向前扫描；
3.如果该元素（已排序）大于新元素，将该元素移到下一位置；
4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
5.将新元素插入到该位置后；
6.重复步骤2~5。
*/
func insertionSort(nums []int) []int {
	var preIndex, current int
	for i := 1; i < len(nums); i++ {
		preIndex, current = i-1, nums[i]
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}
	return nums
}

/*
归并排序

归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

算法描述
1.把长度为n的输入序列分成两个长度为n/2的子序列；
2.对这两个子序列分别采用归并排序；
3.将两个排序好的子序列合并成一个最终的排序序列。
*/
func mergeSort(nums []int) []int {
	var len int = len(nums)
	if len < 2 {
		return nums
	}

	var mid = len >> 1 //类似 mid = len / 2
	left := nums[:mid]
	right := nums[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			//下面两句执行效果类似 result.push(left.shift())
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

/*
快速排序

快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

算法描述
快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

1.从数列中挑出一个元素，称为 “基准”（pivot）；
2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
*/
func quickSort(nums []int, left int, right int) []int {
	var partitionIndex int
	if left < right {
		partitionIndex = partition(nums, left, right)
		quickSort(nums, left, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
	return nums
}

func partition(nums []int, left int, right int) int { //分区操作
	var pivot int = left //基准
	var index int = pivot + 1
	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i] //swap
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot] //swap
	return index - 1
}

/*
需要借助 container/heap 包实现,大顶堆

todo debug
*/
func heapSort(nums []int) []int {
	nums = buildMaxHeap(nums)
	fmt.Println("build:  ", nums)
	len := len(nums)
	for i := len - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] //swap
		len--
		nums = heapify(nums, 0)
		fmt.Println("heapify ", nums)
	}
	return nums
}

/*创建大顶堆*/
func buildMaxHeap(nums []int) []int {
	len := len(nums)
	for i := int(math.Floor(float64(len >> 1))); i >= 0; i-- {
		nums = heapify(nums, i)
	}
	return nums
}

/*堆调整*/
func heapify(nums []int, i int) []int {
	var left, right, largest int = 2*i + 1, 2*i + 2, i
	len := len(nums)
	if left < len && nums[left] > nums[largest] {
		largest = left
	}
	if right < len && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i] //swap
		heapify(nums, largest)
	}
	return nums
}

/*
计数排序
计数排序不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。 作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。

算法描述
1.找出待排序的数组中最大和最小的元素；
2.统计数组中每个值为i的元素出现的次数，存入数组C的第i项；
3.对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）；
4.反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1。
*/
func countingSort(nums []int) []int {
	var maxValue int = 0
	var countMap map[int]int = make(map[int]int)
	var sortedIndex int = 0

	for _, v := range nums {
		countMap[v]++
		if maxValue < v {
			maxValue = v
		}
	}
	//fmt.Println(countMap)
	for j := 0; j < maxValue+1; j++ {
		for countMap[j] > 0 && sortedIndex <= len(nums)-1 {
			nums[sortedIndex] = j //对 nums 重写
			sortedIndex++         //已排序下标往前移动
			countMap[j]--         //计数减一
		}
	}
	return nums
}

/*
桶排序


桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

算法描述
1.设置一个定量的数组当作空桶；
2.遍历输入数据，并且把数据一个一个放到对应的桶里去；
3.对每个不是空的桶进行排序；
4.从不是空的桶里把排好序的数据拼接起来。
*/
func bucketSort(nums []int) []int {
	var max, min int = nums[0], nums[0] //最大最小值,默认是第一个元素
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
	}
	var bucketSize int = 10
	var bucketCount int = int(math.Floor(float64((max-min)/bucketSize))) + 1 //这里一定要向下取整,如 34/10=3.4=>3,进入桶的第 3 个格子
	var buckets [][]int = make([][]int, bucketCount)
	var result []int

	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, bucketCount)
	}
	// fmt.Println(bucketCount)
	for i := 0; i < len(nums); i++ {
		idx := int(math.Floor(float64((nums[i] - min) / bucketSize)))
		//fmt.Println(idx, nums[i])
		buckets[idx] = append(buckets[idx], nums[i])
	}
	fmt.Println(buckets)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = insertionSort(buckets[i]) //用插入排序,将小桶排序
		for j := 0; j < len(buckets[i]); j++ {
			if buckets[i][j] > 0 {
				result = append(result, buckets[i][j])
			}
		}
	}
	return result
}

/*
希尔排序
*/
func shellSort(nums []int) []int {
	length := len(nums)
	gap := 1
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = temp
		}
		gap = gap / 3
	}
	return nums
}

/*
基数排序
基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。有时候有些属性是有优先级顺序的，
先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。

10.1 算法描述
取得数组中的最大数，并取得位数；
arr为原始数组，从最低位开始取每个位组成radix数组；
对radix进行计数排序（利用计数排序适用于小范围数的特点）；

基数排序基于分别排序，分别收集，所以是稳定的。但基数排序的性能比桶排序要略差，每一次关键字的桶分配都需要O(n)的时间复杂度，而且分配之后得到新的关键字序列又需要O(n)的时间复杂度。假如待排数据可以分为d个关键字，则基数排序的时间复杂度将是O(d*2n) ，当然d要远远小于n，因此基本上还是线性级别的。

基数排序的空间复杂度为O(n+k)，其中k为桶的数量。一般来说n>>k，因此额外空间需要大概n个左右
*/
func radixSort(nums []int) []int {
	//获取最大值vl
	vl := 0
	for _, v := range nums {
		if v > vl {
			vl = v
		}
	}
	//获取最大值的位数
	var count int = 0
	for vl%10 > 0 {
		vl = vl / 10
		count++
	}

	//给桶中对应的位置放数据
	for i := 0; i < count; i++ {
		theData := int(math.Pow10(i)) //10的i次方
		//建立空桶
		var bucket [10][10]int
		for k := 0; k < len(nums); k++ {
			theResidue := (nums[k] / theData) % 10 //取余
			var childArray [10]int                 //= bucket[theResidue];//获取子数组
			for m := 0; m < 10; m++ {
				if bucket[theResidue][m] == 0 {
					childArray[m] = nums[k]
					bucket[theResidue][m] = childArray[m]
					break
				} else {
					continue
				}
			}
		}
		//一遍循环完之后需要把数组二维数据进行重新排序，比如数组开始是10 1 18 30 23 12 7 5 18 233 144 ，循环个位数
		//循环之后的结果为10 30 1 12 23 233 144 5 7 18 18 ，然后循环十位数，结果为1 5 7 10 12 18 18 23 30 233 144
		//最后循环百位数，结果为1 5 7 10 12 18 18 23 30 144 233
		var x = 0
		slice := make([]int, len(nums))
		for p := 0; p < len(bucket); p++ {
			for q := 0; q < len(bucket[p]); q++ {
				if bucket[p][q] != 0 {
					slice[x] = bucket[p][q]
					x++
				} else {
					break
				}
			}
		}

		for key, value := range slice {
			nums[key] = value
		}
	}
	return nums

}

// todo TIM 排序
