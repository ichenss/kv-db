package kv_project

import "os"

type Options struct {
	DirPath            string      // 数据库数据目录
	DataFileSize       int64       // 数据文件的大小
	SyncWrites         bool        // 每次写数据是否持久化
	IndexType          IndexerType // 索引类型
	BytesPerSync       uint        // 累计写入多少字节后持久化
	MMapAtStartUp      bool        // 启动时是否使用 MMap 加载数据
	DataFileMergeRatio float32     // 数据文件合并的阈值
}

type IteratorOptions struct {
	// 遍历前缀为指定值的 key，默认为空
	Prefix []byte
	// 是否反向遍历，默认 false 正向
	Reverse bool
}

// WriteBatchOptions 批量写配置项
type WriteBatchOptions struct {
	// 一个批次中最大数据量
	MaxBatchNum uint
	// 提交时是否持久化
	SyncWrites bool
}

type IndexerType = int8

const (
	// Btree 索引
	Btree IndexerType = iota + 1
	// ART 自适应基数树索引
	ART
	// BPlusTree B+ 树索引，将索引存储在磁盘上
	BPlusTree
)

var DefaultOptions = Options{
	DirPath:            os.TempDir(),
	DataFileSize:       256 * 1024 * 1024,
	SyncWrites:         false,
	IndexType:          Btree,
	BytesPerSync:       0,
	MMapAtStartUp:      true,
	DataFileMergeRatio: 0.5,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
