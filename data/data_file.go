package data

import "kv_project/fio"

// DataFile 数据文件
type DataFile struct {
	FileId    uint32        // 文件 id
	WriteOff  int64         // 文件写到了哪个位置
	IOManager fio.IOManager // io 读写管理
}

func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}
