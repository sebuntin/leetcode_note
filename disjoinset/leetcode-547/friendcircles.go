package leetcode_547

// 定义并查集
// 并查集
type DisJoinSet struct {
	id    []int
	size  []int
	Count int
}

// 并查集的构建
func Consturct(max int) DisJoinSet {
	dSet := DisJoinSet{}
	for i := 0; i < max; i++ {
		dSet.id = append(dSet.id, i)
		dSet.size = append(dSet.size, 1)
	}
	dSet.Count = max
	return dSet
}

func (this *DisJoinSet) Find(i int) int {
	for i != this.id[i] {
		this.id[i] = this.id[this.id[i]]
		i = this.id[i]
	}
	return this.id[i]
}

func (this *DisJoinSet) Union(p, q int) {
	i := this.Find(p)
	j := this.Find(q)
	if i == j {
		return
	}
	if this.size[i] > this.size[j] {
		this.id[j] = i
		this.size[i] += this.size[j]
	} else {
		this.id[i] = j
		this.size[j] += this.size[i]
	}
	this.Count--
}

func findCircleNum(M [][]int) int {
	// 构建长度为len(M)的并查集
	dSet := Consturct(len(M))
	for i := 0; i < len(M); i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 {
				dSet.Union(i, j)
			}
		}
	}
	return dSet.Count
}
