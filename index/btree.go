package index

import (
	"bytes"
	"github.com/google/btree"
	"kv_project/data"
	"sort"
	"sync"
)

type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

// NewBTree 初始化 BTree 索引结构体
func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) *data.LogRecordPos {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	oldItem := bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	if oldItem == nil {
		return nil
	}
	return oldItem.(*Item).pos
}

func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}

func (bt *BTree) Delete(key []byte) (*data.LogRecordPos, bool) {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	bt.lock.Unlock()
	return oldItem.(*Item).pos, oldItem != nil
}

func (bt *BTree) Size() int {
	return bt.tree.Len()
}

func (bt *BTree) Iterator(reverse bool) Iterator {
	if bt.tree == nil {
		return nil
	}
	bt.lock.RLock()
	defer bt.lock.RUnlock()
	return newBtreeIterator(bt.tree, reverse)
}

func (bt *BTree) Close() error {
	return nil
}

// B树索引迭代器
type btreeIterator struct {
	currIndex int     // 当前遍历的下标位置
	reverse   bool    // 是否反向遍历
	values    []*Item // key + 位置索引信息
}

func newBtreeIterator(tree *btree.BTree, reverse bool) *btreeIterator {
	var idx int
	values := make([]*Item, tree.Len())

	// 将所有数据存放在数组中
	saveValues := func(it btree.Item) bool {
		values[idx] = it.(*Item)
		idx++
		return true
	}

	if reverse {
		tree.Descend(saveValues)
	} else {
		tree.Ascend(saveValues)
	}

	return &btreeIterator{
		currIndex: 0,
		reverse:   reverse,
		values:    values,
	}
}

// Rewind 重新回到迭代器起点
func (bti *btreeIterator) Rewind() {
	bti.currIndex = 0
}

// Seek 根据 key 查找到第一个大于/小于等于的目标 key，根据从这个 key 开始遍历
func (bti *btreeIterator) Seek(key []byte) {
	if bti.reverse {
		bti.currIndex = sort.Search(len(bti.values), func(i int) bool {
			return bytes.Compare(bti.values[i].key, key) <= 0
		})
	} else {
		bti.currIndex = sort.Search(len(bti.values), func(i int) bool {
			return bytes.Compare(bti.values[i].key, key) >= 0
		})
	}
}

// Next 跳转到下一个 key
func (bti *btreeIterator) Next() {
	bti.currIndex++
}

// Valid 是否有效，即是否已经遍历完了所有的 key，用于退出遍历
func (bti *btreeIterator) Valid() bool {
	return bti.currIndex < len(bti.values)
}

// Key 当前遍历位置的 Key 数据
func (bti *btreeIterator) Key() []byte {
	return bti.values[bti.currIndex].key
}

// Value 当前遍历位置的 Value 数据
func (bti *btreeIterator) Value() *data.LogRecordPos {
	return bti.values[bti.currIndex].pos
}

// Close 关闭迭代器
func (bti *btreeIterator) Close() {
	bti.values = nil
}
