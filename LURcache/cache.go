package LURcache

// 双向链表节点
type node struct {
	key  int
	val  int
	next *node
	prev *node
}

type doubleNodeList struct {
	head *node
	tail *node
}

func newDoubleNodeList() *doubleNodeList {
	head := node{}
	tail := node{prev: &head}
	head.next = &tail
	return &doubleNodeList{
		head: &head,
		tail: &tail,
	}
}

// 从双向链表中删除一个节点
func (this *doubleNodeList) remove(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 往双向链表头部插入节点
func (this *doubleNodeList) headAdd(node *node) {
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
	node.next.prev = node
}

type LRUCache struct {
	cache    map[int]*node
	dList    *doubleNodeList
	capacity int
	size     int
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*node, capacity)
	dList := newDoubleNodeList()
	return LRUCache{
		cache:    cache,
		dList:    dList,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// 将该节点移动到双向链表的头部
	this.dList.remove(node)
	this.dList.headAdd(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node := &node{key: key, val: value}
	// 判断该key值是否已经存在与cache中
	oldNode, ok := this.cache[key]
	if ok {
		this.dList.remove(oldNode)
		this.size --
	}
	if this.size < this.capacity {
		this.dList.headAdd(node)
		this.cache[key] = node
		this.size++
	} else {
		// 先从cache map中删掉尾部节点
		delete(this.cache, this.dList.tail.prev.key)
		// 再从双向链表中将尾部节点删除
		this.dList.remove(this.dList.tail.prev)
		this.dList.headAdd(node)
		this.cache[key] = node
	}
}
