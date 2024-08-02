package fio

const DataFilePerm = 0644

// IOManager 抽象 IO 管理接口，可接入不同 IO 类型，目前支持标准文件 IO
type IOManager interface {
	// Read 从文件给定位置读取对应数据
	Read([]byte, int64) (int, error)

	// Write 写入字节数组到文件中
	Write([]byte) (int, error)

	// Sync 持久化数据
	Sync() error

	// Close 关闭文件
	Close() error

	// Size 获取到文件大小
	Size() (int64, error)
}

func NewIOManager(fileName string) (IOManager, error) {
	return NewFileIOManager(fileName)
}
