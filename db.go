package kv_project

import (
	"kv_project/data"
	"sync"
)

// DB 存储引擎实例
type DB struct {
	mu         *sync.RWMutex
	activeFile *data.DataFile            // 当前活跃数据文件，可用于写入
	olderFile  map[uint32]*data.DataFile // 旧的数据文件，只用于读
}

func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		return ErrKeyIsEmpty
	}

	log_record := data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}
	log_record.Key = []byte("aaa")
	return nil
}

func (db *DB) appendLogRecord() (*data.LogRecordPos, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	return nil, nil
}

func (db *DB) setActiveDataFile() error {
	var initialFileId uint32 = 0
	if db.activeFile != nil {
		initialFileId = db.activeFile.FileId + 1
	}

	// 打开新的数据文件
	data.OpenDataFile()
	return nil
}
