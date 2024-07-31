package kv_project

type Options struct {
	DirPath      string // 数据库数据目录
	DataFileSize int64  // 数据文件的大小
	SyncWrites   bool   // 每次写数据是否持久化
}
