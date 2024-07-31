package data

import "kv_project/fio"

// DataFile 数据文件
type DataFile struct {
	FileId    uint32        // 文件 id
	WriteOff  int64         // 文件写到了哪个位置
	IOManager fio.IOManager // io 读写管理
}

// OpenDataFile 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {

	return nil, nil
}

func (df *DataFile) ReadLogRecord(offset int64) (*LogRecord, error) {
	return nil, nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) Sync() error {
	return nil
}