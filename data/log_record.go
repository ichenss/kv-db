package data

import (
	"encoding/binary"
	"hash/crc32"
)

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
	LogRecordTxnFinished
)

const maxLogRecordHeaderSize = binary.MaxVarintLen32*2 + 5

// LogRecord 写入到数据文件的记录
// 之所以叫日志，数据文件中的数据是追加写入，类似日志格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// logRecordHeader LogRecord 的头部信息
type logRecordHeader struct {
	crc        uint32        // crc 校验值
	recordType LogRecordType // 标识 LogRecord 类型
	keySize    uint32        // key 的长度
	valueSize  uint32        // value 的长度
}

// LogRecordPos 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id，表示将文件存储到哪个文件当中
	Offset int64  // 偏移，表示将数据存储到数据文件中哪个位置
}

// EncodeLogRecord 对 logRecord 进行编码，返回字节数组及长度
// crc | type | keyIndex | valueIndex | Key | Value
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	header := make([]byte, maxLogRecordHeaderSize)

	header[4] = logRecord.Type

	var index = 5

	index += binary.PutVarint(header[index:], int64(len(logRecord.Key)))
	index += binary.PutVarint(header[index:], int64(len(logRecord.Value)))

	size := len(logRecord.Key) + len(logRecord.Value) + index
	encBytes := make([]byte, size)

	copy(encBytes[:index], header[:index])

	copy(encBytes[index:], logRecord.Key)
	copy(encBytes[index+len(logRecord.Key):], logRecord.Value)

	crc := crc32.ChecksumIEEE(encBytes[4:])
	binary.LittleEndian.PutUint32(encBytes[:4], crc)

	return encBytes, int64(size)
}

// decodeLogRecordHeader 对字节数组中的 Header 信息进行解码
func decodeLogRecordHeader(buf []byte) (*logRecordHeader, int64) {
	if len(buf) < 4 {
		return nil, 0
	}

	header := &logRecordHeader{
		crc:        binary.LittleEndian.Uint32(buf[:4]),
		recordType: buf[4],
	}

	var index = 5
	keySize, n := binary.Varint(buf[index:])
	index += n
	valueSize, n := binary.Varint(buf[index:])
	index += n

	header.keySize = uint32(keySize)
	header.valueSize = uint32(valueSize)

	return header, int64(index)
}

// getLogRecordCRC
func getLogRecordCRC(lr *LogRecord, header []byte) uint32 {
	if lr == nil {
		return 0
	}

	crc := crc32.ChecksumIEEE(header[:])
	crc = crc32.Update(crc, crc32.IEEETable, lr.Key)
	crc = crc32.Update(crc, crc32.IEEETable, lr.Value)

	return crc
}
