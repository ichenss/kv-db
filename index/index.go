package index

import (
	"bytes"
	"github.com/google/btree"
	"kv_project/data"
)

// Indexer 抽象索引接口，后续如果想要实现其他数据结构，实现这个接口即可
type Indexer interface {
	// Put 像索引中存储 key 对应的位置信息
	Put(key []byte, pos *data.LogRecordPos) bool

	// Get 根据 key 取出对应位置索引信息
	Get(key []byte) *data.LogRecordPos

	// Delete 根据 key 删除对应位置索引信息
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}
