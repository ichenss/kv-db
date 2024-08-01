package kv_project

type Options struct {
	DirPath      string      // 数据库数据目录
	DataFileSize int64       // 数据文件的大小
	SyncWrites   bool        // 每次写数据是否持久化
	IndexType    IndexerType // 索引类型
}

type IndexerType = int8

const (
	// Btree 索引
	Btree IndexerType = iota + 1
	// ART 自适应基数树索引
	ART
)
