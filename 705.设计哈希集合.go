package leetcode

/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
const base int = 679

type MyHashSet struct {
	hash []*Listnode
}

type Listnode struct {
	Key  int
	Next *Listnode
}

func Constructor() MyHashSet {
	return MyHashSet{hash: make([]*Listnode, base)}
}

func (this *MyHashSet) Add(key int) {
	index := key % base
	node := this.hash[index]
	if node == nil {
		this.hash[index] = &Listnode{Key: key}
	} else {
		if node.Key == key {
			return
		}
		for node.Next != nil {
			if node.Next.Key == key {
				return
			}
			node = node.Next
		}
		node.Next = &Listnode{Key: key}
	}
}

func (this *MyHashSet) Remove(key int) {
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

func (this *MyHashSet) Contains(key int) bool {
	index := key % base
	node := this.hash[index]
	if node == nil {
		return false
	} else {
		if node.Key == key {
			return true
		}
		for node.Next != nil {
			if node.Next.Key == key {
				return true
			}
			node = node.Next
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end
