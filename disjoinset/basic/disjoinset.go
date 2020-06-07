package main

import "fmt"

// 并查集
type Disjoinset struct {
	id    []int // 表示原数组
	size  []int // 表示集合的size
	Count int
}

// 并查集的构建
func Consturct(max int) Disjoinset {
	// 初始化并查集的时候将原数组的每个元素都作为一个集合，
	//这样每个集合的size都为1
	dSet := Disjoinset{}
	for i := 0; i < max; i++ {
		dSet.id = append(dSet.id, i)
		dSet.size = append(dSet.size, 1)
	}
	dSet.Count = max
	return dSet
}

func (this *Disjoinset) Find(i int) int {
	for i != this.id[i] {
		this.id[i] = this.id[this.id[i]]
		i = this.id[i]
	}
	return this.id[i]
}

func (this *Disjoinset) Union(p, q int) {
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

func main() {
	dSet := Consturct(8)
	fmt.Println(dSet.Find(1))
	dSet.Union(0, 1)
	dSet.Union(4, 0)
	fmt.Println(dSet.Find(0))
	fmt.Println(dSet.Find(1))
	fmt.Println(dSet.Find(4))
	fmt.Println(dSet.Count)
}
