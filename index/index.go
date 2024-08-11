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

	// Size 索引中的数据量
	Size() int

	// Iterator 索引迭代器
	Iterator(reverse bool) Iterator
}

type IndexType = int8

const (
	// Btree 索引
	Btree IndexType = iota + 1
	// ART 自适应基数树索引
	ART
)

func NewIndexer(typ IndexType) Indexer {
	switch typ {
	case Btree:
		return NewBTree()
	case ART:
		// todo
		return nil
	default:
		panic("unsupported index type")
	}
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

// Iterator 通用索引迭代器
type Iterator interface {
	// Rewind 重新回到迭代器起点
	Rewind()
	// Seek 根据 key 查找到第一个大于/小于等于的目标 key，根据从这个 key 开始遍历
	Seek(key []byte)
	// Next 跳转到下一个 key
	Next()
	// Valid 是否有效，即是否已经遍历完了所有的 key，用于推出遍历
	Valid() bool
	// Key 当前遍历位置的 Key 数据
	Key() []byte
	// Value 当前遍历位置的 Value 数据
	Value() *data.LogRecordPos
	// Close 关闭迭代器
	Close()
}
