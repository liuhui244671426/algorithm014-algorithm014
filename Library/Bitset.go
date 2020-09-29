package Library

import "math/bits"

/*
文档讲解出处: https://zhuanlan.zhihu.com/p/90749142
github 出处: https://github.com/poloxue/go-set-example
*/
const (
	shift = 6    // 2^n = 64 的n
	mask  = 0x3f //n=6,2^n-1=63=0x3f
)

type BitSet struct {
	data []uint64 //动态可扩容
	size int      //元素个数无法通过 len 获取
}

/*
假设用 index 表示行（索引），pos 表示列（位置）
index 可通过元素值整除字长，即 value / 64，转化为高效的位运算，即 value >> 6
pos 可以通过元素值取模字长，即 value % 64，转化为高效的位运算，即 value & 0x3f，获取对应位置，
然后用 1 << uint(value % 0xf) 即可将位置转化为值
*/
func index(n int) int {
	return n >> shift
}

// 相对于标志位使用场景中某个标志的值
func posValue(n int) uint64 {
	return 1 << uint(n&mask)
}

func NewBitset(ns ...int) *BitSet {
	//如果输入参数 ns 为空的话，new(BitSet) 返回空集合即可。
	if len(ns) == 0 {
		return new(BitSet)
	}
	// 计算多 bitset 开辟多个空间
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	// 如果 max < 0，直接返回空。
	if max < 0 {
		return new(BitSet)
	}
	// 通过 max >> shift+1 计算最大值 max 所在 index
	// 而 index + 1 即为要开辟的空间
	s := &BitSet{
		data: make([]uint64, index(max)+1),
	}

	for _, n := range ns {
		if n >= 0 {
			// e >> shift 获取索引位置，即行，一般叫 index
			// e&mask 获取所在列，一般叫 pos，F1 0 F2 1
			s.data[n>>shift] |= posValue(n)
			//添加元素个数
			s.size++
		}
	}
	return s
}

func (set *BitSet) Add(n int) *BitSet {
	if n < 0 {
		return set
	}
	// 检测是否有足够的空间存放新元素
	i := index(n)
	if i >= len(set.data) {
		ndata := make([]uint64, i+1)
		copy(ndata, set.data)
		set.data = ndata
	}

	if set.data[i]&posValue(n) == 0 {
		// 设置元素到集合中
		set.data[i] |= posValue(n)
		set.size++
	}
	return set
}

func (set *BitSet) Clear(n int) *BitSet {
	if n < 0 {
		return set
	}
	i := index(n)
	if i >= len(set.data) {
		return set
	}
	if set.data[i]&posValue(n) != 0 {
		//通过 &^ 实现指定位清除
		set.data[i] &^= posValue(n)
		set.size--
	}
	return set
}

// 获取元素对应的 int64 的位置，如果超出 data 能表示的范围，直接返回。
func (set *BitSet) Contains(n int) bool {
	i := index(n)
	if i >= len(set.data) {
		return false
	}
	return set.data[i]&posValue(n) != 0
}

//计算集合中的元素个数
func (set *BitSet) computeSize() int {
	d := set.data
	n := 0
	for i, l := 0, len(d); i < l; i++ {
		if w := d[i]; w != 0 {
			n += bits.OnesCount64(w)
		}
	}
	return n
}

// 交集
func (set *BitSet) Intersect(other *BitSet) *BitSet {
	minLen := min(len(set.data), len(other.data))
	intersectSet := &BitSet{
		data: make([]uint64, minLen),
	}

	for i := 0; i < minLen; i++ {
		intersectSet.data[i] = set.data[i] & other.data[i]
	}
	intersectSet.size = set.computeSize()
	return intersectSet
}

// 并集
func (set *BitSet) Union(other *BitSet) *BitSet {
	var maxSet, minSet *BitSet
	if len(set.data) > len(other.data) {
		maxSet, minSet = set, other
	} else {
		maxSet, minSet = other, set
	}

	unionSet := &BitSet{
		data: make([]uint64, len(maxSet.data)),
	}
	minLen := len(minSet.data)
	copy(unionSet.data[minLen:], maxSet.data[minLen:])

	for i := 0; i < minLen; i++ {
		unionSet.data[i] = set.data[i] | other.data[i]
	}
	unionSet.size = unionSet.computeSize()
	return unionSet
}

// 差集
func (set *BitSet) Difference(other *BitSet) *BitSet {
	setLen := len(set.data)
	otherLen := len(other.data)

	differenceSet := &BitSet{
		data: make([]uint64, setLen),
	}

	minLen := setLen
	if setLen > otherLen {
		copy(differenceSet.data[otherLen:], set.data[otherLen:])
		minLen = otherLen
	}
	for i := 0; i < minLen; i++ {
		differenceSet.data[i] = set.data[i] &^ other.data[i]
	}
	differenceSet.size = differenceSet.computeSize()
	return differenceSet
}
func (set *BitSet) Visit(do func(int) (skip bool)) (aborted bool) {
	d := set.data
	for i, len := 0, len(d); i < len; i++ {
		w := d[i]
		if w == 0 {
			continue
		}
		// 0 << 6，还是 0，1 << 6 就是 64，2 << 6 的就是 128
		n := i << shift
		for w != 0 {
			// 000.....000100 64~128 的话，表示 66，即 64 + 2，这个 2 可以由结尾 0 的个数确定
			// 那怎么获取结果 0 的个数呢？可以使用 bits.TrailingZeros64 函数
			b := bits.TrailingZeros64(w)
			if do(n + b) {
				return true
			}
			// 将已经检查的位清零
			// 为了保证尾部 0 的个数能代表元素的值
			w &^= 1 << uint64(b)
		}
	}
	return false
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func (set *BitSet) Size() int {
	return set.size
}

//等同于 reslice ,用来收缩 data 空间
func (set *BitSet) trim() {
	d := set.data
	n := len(d) - 1
	for n >= 0 && d[n] == 0 {
		n--
	}
	set.data = d[:n+1]
}
