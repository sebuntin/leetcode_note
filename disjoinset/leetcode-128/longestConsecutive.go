package leetcode_128

type DisJoinSet struct {
	id   map[int]int
	size map[int]int
}

func Construct(nums []int) *DisJoinSet {
	dset := &DisJoinSet{
		id:   make(map[int]int),
		size: make(map[int]int),
	}
	for _, num := range nums {
		dset.id[num] = num
		dset.size[num] = 1
	}
	return dset
}

func (this *DisJoinSet) Find(i int) (int, bool) {
	if _, ok := this.id[i]; !ok {
		return 0, ok
	} else {
		for i == this.id[i] {
			this.id[i] = this.id[this.id[i]]
		}
	}
	return this.id[i], true
}

func (this *DisJoinSet) Union(i, j int) {
	id_i, ok := this.Find(i)
	if !ok {
		return
	}
	id_j, ok := this.Find(j)
	if !ok {
		return
	}
	if id_i == id_j {
		return
	}
	if this.size[id_i] < this.size[id_j] {
		this.size[id_j] += this.size[id_i]
		this.id[i] = id_j
		return
	} else {
		this.size[id_i] += this.size[id_j]
		this.id[j] = id_i
		return
	}

}

func longestConsecutive(nums []int) int {
	// 构建该数组的并查集
	dset := Construct(nums)
	// 将num与num+1所属集合进行union操作
	maxLen := 1
	for i := range nums {
		if _, ok := dset.Find(nums[i] + 1); ok {
			dset.Union(nums[i], nums[i]+1)
		}
		idx, _ := dset.Find(nums[i])
		maxLen = Max(dset.size[idx], maxLen)
	}
	return maxLen
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
