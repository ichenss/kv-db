package index

import (
	"go.etcd.io/bbolt"
	"kv_project/data"
)

// BPlusTree B+ 树索引
type BPlusTree struct {
	tree bbolt.DB
}

// NewBPlusTree 初始化 B+ 树索引
func NewBPlusTree() *BPlusTree {
	return nil
}

// Put 像索引中存储 key 对应的位置信息
func (bpt *BPlusTree) Put(key []byte, pos *data.LogRecordPos) bool {
	return false
}

// Get 根据 key 取出对应位置索引信息
func (bpt *BPlusTree) Get(key []byte) *data.LogRecordPos {
	return nil
}

// Delete 根据 key 删除对应位置索引信息
func (bpt *BPlusTree) Delete(key []byte) bool {
	return false
}

// Size 索引中的数据量
func (bpt *BPlusTree) Size() int {
	return 0
}

// Iterator 索引迭代器
func (bpt *BPlusTree) Iterator(reverse bool) Iterator {
	return nil
}
