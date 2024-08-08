package index

import (
	"github.com/stretchr/testify/assert"
	"kv_project/data"
	"testing"
)

func TestBTree_Put(t *testing.T) {
	bt := NewBTree()

	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)

	res2 := bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)
}

func TestBTree_Get(t *testing.T) {
	bt := NewBTree()

	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)

	pos1 := bt.Get(nil)
	assert.Equal(t, uint32(1), pos1.Fid)
	assert.Equal(t, int64(100), pos1.Offset)

	res2 := bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)

	res3 := bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 3})
	assert.True(t, res3)

	pos2 := bt.Get([]byte("a"))
	t.Log(pos2)
	assert.Equal(t, uint32(1), pos2.Fid)
	assert.Equal(t, int64(3), pos2.Offset)
}

func TestBTree_Delete(t *testing.T) {
	bt := NewBTree()
	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)

	res2 := bt.Delete(nil)
	assert.True(t, res2)

	res3 := bt.Put([]byte("aaa"), &data.LogRecordPos{Fid: 22, Offset: 333})
	assert.True(t, res3)

	res4 := bt.Delete([]byte("aaa"))
	assert.True(t, res4)

}

func TestBTree_Iterator(t *testing.T) {
	bt1 := NewBTree()
	// 1. BTree 为空情况
	it1 := bt1.Iterator(false)
	assert.Equal(t, false, it1.Valid())

	// 2. 有数据情况
	bt1.Put([]byte("code"), &data.LogRecordPos{
		Fid:    1,
		Offset: 10,
	})
	it2 := bt1.Iterator(false)
	assert.Equal(t, true, it2.Valid())
	assert.NotNil(t, it2.Key())
	assert.NotNil(t, it2.Value())
	it2.Next()
	assert.Equal(t, false, it2.Valid())

	// 3. 多条数据
	bt1.Put([]byte("b"), &data.LogRecordPos{
		Fid:    1,
		Offset: 10,
	})
	bt1.Put([]byte("a"), &data.LogRecordPos{
		Fid:    1,
		Offset: 10,
	})
	bt1.Put([]byte("c"), &data.LogRecordPos{
		Fid:    1,
		Offset: 10,
	})
	it3 := bt1.Iterator(false)
	for it3.Rewind(); it3.Valid(); it3.Next() {
		assert.NotNil(t, it3.Key())
	}

	it4 := bt1.Iterator(true)
	for it4.Rewind(); it4.Valid(); it4.Next() {
		assert.NotNil(t, it4.Key())
	}

	// 4. seek 测试
	it5 := bt1.Iterator(false)
	it5.Seek([]byte("b"))
	for it5.Valid() {
		assert.NotNil(t, it5.Key())
		it5.Next()
	}

	it6 := bt1.Iterator(true)
	it6.Seek([]byte("c"))
	for it6.Valid() {
		t.Log(string(it6.Key()))
		assert.NotNil(t, it6.Key())
		it6.Next()
	}
}
