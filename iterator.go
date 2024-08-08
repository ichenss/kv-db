package kv_project

import (
	"bytes"
	"kv_project/index"
)

type Iterator struct {
	indexIter index.Iterator
	db        *DB
	options   IteratorOptions
}

// NewIterator 初始化迭代器
func (db *DB) NewIterator(opts IteratorOptions) *Iterator {
	indexIter := db.index.Iterator(opts.Reverse)
	return &Iterator{
		indexIter: indexIter,
		db:        db,
		options:   opts,
	}
}

// Rewind 重新回到迭代器起点
func (it *Iterator) Rewind() {
	it.indexIter.Rewind()
	it.skipToNext()
}

// Seek 根据 key 查找到第一个大于/小于等于的目标 key，根据从这个 key 开始遍历
func (it *Iterator) Seek(key []byte) {
	it.indexIter.Seek(key)
	it.skipToNext()
}

// Next 跳转到下一个 key
func (it *Iterator) Next() {
	it.indexIter.Next()
	it.skipToNext()
}

// Valid 是否有效，即是否已经遍历完了所有的 key，用于推出遍历
func (it *Iterator) Valid() bool {
	return it.indexIter.Valid()
}

// Key 当前遍历位置的 Key 数据
func (it *Iterator) Key() []byte {
	return it.indexIter.Key()
}

// Value 当前遍历位置的 Value 数据
func (it *Iterator) Value() ([]byte, error) {
	pos := it.indexIter.Value()
	it.db.mu.RLock()
	defer it.db.mu.RUnlock()
	return it.db.getValueByPosition(pos)
}

// Close 关闭迭代器
func (it *Iterator) Close() {
	it.indexIter.Close()
}

func (it *Iterator) skipToNext() {
	prefixNum := len(it.options.Prefix)
	if prefixNum == 0 {
		return
	}

	for ; it.indexIter.Valid(); it.indexIter.Next() {
		key := it.indexIter.Key()
		if prefixNum <= len(key) && bytes.Compare(it.options.Prefix, key[:prefixNum]) == 0 {
			break
		}
	}
}
