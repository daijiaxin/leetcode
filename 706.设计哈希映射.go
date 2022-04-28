package leetcode

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
const base int = 679

type MyHashMap struct {
	hash []*Listnode
}

type Listnode struct {
	Key   int
	Value int
	Next  *Listnode
}

func Constructor() MyHashMap {
	return MyHashMap{hash: make([]*Listnode, base)}
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % base
	node := this.hash[index]
	if node == nil {
		this.hash[index] = &Listnode{Key: key, Value: value}
	} else {
		if node.Key == key {
			node.Value = value
			return
		}
		for node.Next != nil {
			if node.Next.Key == key {
				node.Next.Value = value
				return
			}
			node = node.Next
		}
		node.Next = &Listnode{Key: key, Value: value}
	}
}

func (this *MyHashMap) Get(key int) int {
	index := key % base
	node := this.hash[index]
	if node == nil {
		return -1
	} else {
		if node.Key == key {
			return node.Value
		}
		for node.Next != nil {
			if node.Next.Key == key {
				return node.Next.Value
			}
			node = node.Next
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % base
	node := this.hash[index]
	if node != nil {
		if node.Key == key {
			this.hash[index] = node.Next
		} else {
			for node.Next != nil {
				if node.Next.Key == key {
					node.Next = node.Next.Next
					return
				}
				node = node.Next
			}
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
